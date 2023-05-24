package main

import (
	"net/http"
	"os"
	"testing"
)

type myHandler struct{}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
