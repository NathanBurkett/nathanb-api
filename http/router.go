package http

import (
	"net/http"
	"github.com/nathanburkett/nathanb-api/app"
	"time"
	"log"
	)

type Router struct {
	Server *http.ServeMux
	Routes []Route
	app    *app.Instance
}

func NewRouter(instance *app.Instance) *Router {
	return &Router{
		Server: http.NewServeMux(),
		Routes: []Route{},
		app:    instance,
	}
}

func (r *Router) attachRoute(route Route) *Router {
	r.Routes = append(r.Routes, route)

	return r
}

func (r *Router) attachRouteTree(tree RouteTree) *Router {
	routes := tree.Define(r.app)

	for i := 0; i < len(routes); i++ {
		r.attachRoute(routes[i])
	}

	return r
}

func (r *Router) Attach(addtl ...RouteTree) *Router {
	trees := r.getDefaultRouteTrees()

	for i := 0; i < len(addtl); i++ {
		trees = append(trees, addtl[i])
	}

	for i := 0; i < len(trees); i++ {
		r.attachRouteTree(trees[i])
	}

	return r
}

func (r *Router) Define(route Route) {
	r.Server.Handle(route.Pattern, r.apply(route))
}

func (r *Router) apply(route Route) http.Handler {
	if route.Middleware == nil {
		return route
	}

	var handler http.Handler
	handler = route

	for i := len(route.Middleware) - 1; i >= 0; i-- {
		handler = route.Middleware[i](handler)

	}

	return handler
}

func (r *Router) Mount() *Router {
	for i := 0; i < len(r.Routes); i++ {
		r.Define(r.Routes[i])
	}

	return r
}

func (r *Router) getDefaultRouteTrees() []RouteTree {
	return []RouteTree{
		graphqlRouteTree{},
	}
}

func (r *Router) ListenAndServe(addr string) error {
	s := &http.Server{
		Addr: addr,
		Handler: r.Server,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 90 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	log.Printf("Listening for requests on %s", s.Addr)

	return s.ListenAndServe()
}
