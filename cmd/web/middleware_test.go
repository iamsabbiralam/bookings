package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myH myHandler
	h := NoSurf(&myH)
	switch v := h.(type) {
		case http.Handler:
			// do nothing; test passed
		default:
			t.Error(fmt.Sprintf("type is not http.Handler but it %T", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler
	h := SessionLoad(&myH)
	switch v := h.(type) {
		case http.Handler:
			// do nothing; test passed
		default:
			t.Error(fmt.Sprintf("type is not http.Handler but it %T", v))
	}
}
