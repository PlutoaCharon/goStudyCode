package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

func (f HandlerFunc) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

func HelloHandler(res http.ResponseWriter, req *http.Request) {
	_, _ = res.Write([]byte("Hello world"))
}

func main() {
	hf := HandlerFunc(HelloHandler)

	resq := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", bytes.NewBuffer([]byte("test")))

	hf.ServerHTTP(resq, req)
	bts, _ := ioutil.ReadAll(resq.Body)
	fmt.Println(string(bts))
}
