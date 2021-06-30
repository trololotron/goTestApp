package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Hello(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fullName := vars["fullName"]
	fmt.Fprint(response, "Hello "+fullName)
}

func add(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	a := vars["a"]
	b := vars["b"]
	n1, _ := strconv.ParseInt(a, 10, 64)
	n2, _ := strconv.ParseInt(b, 10, 64)
	fmt.Fprint(response, n1 + n2)
}

func sub(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	a := vars["a"]
	b := vars["b"]
	n1, _ := strconv.ParseInt(a, 10, 64)
	n2, _ := strconv.ParseInt(b, 10, 64)
	fmt.Fprint(response, n1 - n2)
}

func mul(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	a := vars["a"]
	b := vars["b"]
	n1, _ := strconv.ParseInt(a, 10, 64)
	n2, _ := strconv.ParseInt(b, 10, 64)
	fmt.Fprint(response, n1 * n2)
}

func div(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	a := vars["a"]
	b := vars["b"]
	n1, _ := strconv.ParseInt(a, 10, 64)
	n2, _ := strconv.ParseInt(b, 10, 64)
	fmt.Fprint(response, n1 / n2)
}