package http

import (
	"net/http"
	"bytes"
	"fmt"
)

type Response interface {
	Code() int
	SetCode(int)
	Body() []byte
	SetBody([]byte)
	Write(w http.ResponseWriter)
}

type JsonResponse struct {
	code  int
	body  []byte
}

func NewJsonResponse() *JsonResponse {
	return &JsonResponse{
		code: 200,
		body: []byte(""),
	}
}

func (json *JsonResponse) Code() int {
	return json.code
}

func (json *JsonResponse) SetCode(code int)  {
	json.code = code
}

func (json *JsonResponse) Body() []byte {
	return json.body
}

func (json *JsonResponse) SetBody(body []byte) {
	json.body = body
}

func (json *JsonResponse) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(json.Code())
	_, _ = w.Write(json.Body())
}

func errorJSON(msg string) []byte {
	buf := bytes.Buffer{}
	fmt.Fprintf(&buf, `{"error": "%s"}`, msg)
	return buf.Bytes()
}
