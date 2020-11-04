package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":3000", nil)
}

type Resp struct {
	Message string
}

func hello(w http.ResponseWriter, req *http.Request) {
	resp := Resp{"Hello world !!"}
	js, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

//https://www.alexedwards.net/blog/golang-response-snippets