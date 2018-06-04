package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/diskevent", diskAddHandler).Methods("POST")
	router.HandleFunc("/diskevent", diskRemoveHandler).Methods("DELETE")
	err := http.ListenAndServe(":8080", router)
	if err!= nil{
	    fmt.Println(err)
	}
}

func diskAddHandler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	fmt.Println(r.Method)
	fmt.Println(string(bodyBytes))
	w.WriteHeader(http.StatusCreated)
	return
}

func diskRemoveHandler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	fmt.Println(r.Method)
	fmt.Println(string(bodyBytes))
	w.WriteHeader(http.StatusCreated)
	return
}
