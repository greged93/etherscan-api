package main

import (
	"net/http"
)

type HttpRequest struct {
	url          string
	rawResponse  *http.Response
	jsonResponse map[string]interface{}
	err          error
}

func (h *HttpRequest) makeRequest(url string) {
	h.rawResponse, h.err = http.Get(url)
}

func main() {
	var req HttpRequest
	req.makeRequest("https://jsonplaceholder.typicode.com/posts/1")
}
