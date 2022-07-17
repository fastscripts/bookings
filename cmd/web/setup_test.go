package main

import (
	"net/http"
	"os"
	"testing"
)

func TestMain(t *testing.M) {

	os.Exit(t.Run())
}

type myHandler struct{}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

/*
go test -coverprofile=coverage.out && go tool cover -html=coverage.out
*/
