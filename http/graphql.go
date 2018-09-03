package http

import (
	"github.com/nathanburkett/graphql-go"
	gqlErrors "github.com/nathanburkett/graphql-go/errors"
	"github.com/nathanburkett/nathanb-api/app"
	"net/http"
	"encoding/json"
	"sync"
)

type graphqlRouteTree struct {}

func (gql graphqlRouteTree) Define(instance *app.Instance) []Route {
	var routes []Route
	routes = append(routes, NewRoute("/graphql", instance, graphqlHandler))
	return routes
}

func graphqlHandler(i *app.Instance, w http.ResponseWriter, r *http.Request) {
	res := NewJsonResponse()

	req, err := parse(r)
	if err != nil {
		res.SetCode(http.StatusBadRequest)
		res.SetBody([]byte(err.Error()))
		res.Write(w)
		return
	}

	n := len(req.queries)
	if n == 0 {
		res.SetCode(http.StatusBadRequest)
		res.SetBody([]byte("no queries to execute"))
		res.Write(w)
		return
	}

	var responses = make([]*graphql.Response, n)
	var wg        sync.WaitGroup

	wg.Add(n)

	for index, q := range req.queries {
		go func(index int, q query) {
			tempRes := i.Schema.Exec(r.Context(), q.Query, q.OpName, q.Variables)

			tempRes.Errors = expand(tempRes.Errors)

			// We have to do some work here to expand.go errors when it is possible for a resolver to return
			// more than one error (for example, a list resolver).
			//res.Errors = errors.Expand(res.Errors)

			responses[index] = tempRes
			wg.Done()
		}(index, q)
	}

	wg.Wait()

	var resp []byte
	if req.isBatch {
		resp, err = json.Marshal(responses)
	} else if len(responses) > 0 {
		resp, err = json.Marshal(responses[0])
	}

	if err != nil {
		res.SetCode(http.StatusInternalServerError)
		res.SetBody(errorJSON("server error"))
		res.Write(w)
		return
	}

	res.SetBody(resp)
	res.Write(w)
	return
}

type slicer interface {
	Slice() []error
}

type indexedCauser interface {
	Index() int
	Cause() error
}

func expand(errs []*gqlErrors.QueryError) []*gqlErrors.QueryError {
	expanded := make([]*gqlErrors.QueryError, 0, len(errs))

	for _, err := range errs {
		switch t := err.ResolverError.(type) {
		case slicer:
			for _, e := range t.Slice() {
				qe := &gqlErrors.QueryError{
					Message:   err.Message,
					Locations: err.Locations,
					Path:      err.Path,
				}

				if ic, ok := e.(indexedCauser); ok {
					qe.Path = append(qe.Path, ic.Index())
					qe.Message = ic.Cause().Error()
				}

				expanded = append(expanded, qe)
			}
		default:
			expanded = append(expanded, err)
		}
	}

	return expanded
}
