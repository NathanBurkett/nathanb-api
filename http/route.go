package http

import (
	"github.com/nathanburkett/nathanb-api/app"
	"net/http"
	"fmt"
	)

const RequestKey = "request"

type RouteHandler func(*app.Instance, http.ResponseWriter, *http.Request)

type Route struct {
	Pattern          string
	App              *app.Instance
	Handler          RouteHandler
	Middleware       []Middleware
}

func NewRoute(pattern string, app *app.Instance, handler RouteHandler, middleware ...Middleware) Route {
	return Route{
		Pattern:    pattern,
		App:        app,
		Handler:    handler,
		Middleware: middleware,
	}
}

type RouteTree interface {
	Define(*app.Instance) []Route
}

func (r Route) addContext(req *http.Request, fn ContextualizeRequest) *http.Request {
	return req.WithContext(fn(req.Context()))
}

func (r Route) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	r.Handler(r.App, writer, req)
}

func (r Route) Error(w http.ResponseWriter, error string, code int) {
	fmt.Fprintln(w, error)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
}
