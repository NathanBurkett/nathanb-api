package http

import (
	"net/http"
	"github.com/nathanburkett/nathanb-api/service"
	"log"
)

type Middleware func (route http.Handler) http.Handler

func AddDbContext (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v", r)

		next.ServeHTTP(w, r)
	})
}

func AddRequestLogger (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s", r.URL.Path)

		next.ServeHTTP(w, r)
	})
}

func AdminMiddleware (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := service.NewAuthService()
		// get user repository
		// find user from token

		if !auth.IsAuthenticated() {
			// handle
		}

		next.ServeHTTP(w, r)
	})
}

func AttachQuerySchema (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		switch r.URL.Path {

		}

		// do stuff
		next.ServeHTTP(w, r)
	})
}

func AttachMutationSchema (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		switch r.URL.Path {

		}

		// do stuff
		next.ServeHTTP(w, r)
	})
}
