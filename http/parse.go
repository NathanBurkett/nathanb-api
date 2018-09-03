package http

import (
	"net/http"
	"io/ioutil"
	"errors"
	"net/url"
	"encoding/json"
	)

// A request respresents an HTTP request to the GraphQL endpoint.
// A request can have a single query or a batch of requests with one or more queries.
// It is important to distinguish between a single query request and a batch request with a single query.
// The shape of the response will differ in both cases.
type request struct {
	queries []query
	isBatch bool
}

// A query represents a single GraphQL query.
type query struct {
	OpName    string                 `json:"operationName"`
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

func parse(r *http.Request) (request, error) {
	// We always need to read and close the request body.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return request{}, errors.New("unable to read request body")
	}
	_ = r.Body.Close()

	var req request

	switch r.Method {
	case "POST":
		req = parsePost(body)
	case "GET":
		req = parsePost(body)
	default:
		err = errors.New("only POST and GET requests are supported")
	}

	return req, err
}

func parseGet(values url.Values) request {
	var (
		queries   = values["query"]
		names     = values["operationName"]
		variables = values["variables"]
		qLen      = len(queries)
		nLen      = len(names)
		vLen      = len(variables)
	)

	if qLen == 0 {
		return request{}
	}

	var requests = make([]query, 0, qLen)
	var isBatch bool

	// This loop assumes there will be a corresponding element at each index
	// for query, operation name, and variable fields.
	//
	// NOTE: This could be a bad assumption. Maybe we want to do some validation?
	for i := 0; i < qLen; i++ {
		q := queries[i]

		var n string
		if i < nLen {
			n = names[i]
		}

		var m = map[string]interface{}{}
		if i < vLen {
			str := variables[i]
			if err := json.Unmarshal([]byte(str), &m); err != nil {
				m = nil // TODO: Improve error handling here.
			}
		}

		requests = append(requests, query{Query: q, OpName: n, Variables: m})
	}

	if qLen > 1 {
		isBatch = true
	}

	return request{queries: requests, isBatch: isBatch}
}

func parsePost(b []byte) request {
	if len(b) == 0 {
		return request{}
	}

	var queries []query
	var isBatch bool

	switch b[0] {
	case '{':
		q := query{}
		if err := json.Unmarshal(b, &q); err == nil {
			queries = append(queries, q)
		}
	case '[':
		isBatch = true
		_ = json.Unmarshal(b, &queries)
	}

	return request{queries: queries, isBatch: isBatch}
}
