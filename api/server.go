package main

import {
	"fmt"
	"net/http"
}

const PORT = "8080"

type Handler = func(http.ResponseWriter, r *http.Request) error

func (e route) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	funcHandler, ok := e.method[r.Method]

	if !ok {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	err := funcHandler(w, r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func Start() {
	fmt.Println("Starting Onigiri Go Server on port %s\n", PORT)
	
	for _, route := range routeSchema {
		http.HandleFunc(route.name, route.ServeHTTP)
	}

	http.ListenAndServe(":"+PORT, nil)
}
