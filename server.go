package main

import (
	"log"
	"net/http"
	"encoding/json"
	"fmt"
)

type Zurl struct {
	Id string
	LongUrl string
}

func main() {
	http.HandleFunc("/", Root)

	fmt.Println("listening...")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("Error: %v", err)
	}
}

func Root(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	switch req.Method {
	case "GET":
		Expand(res, req)
	case "POST":
		Shorten(res, req)
	default:
		res.WriteHeader(405)
	}
}

func Expand(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(res, "Hello! Zurl is a URL shortener service.")
}

func Shorten(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	url := &Zurl{Id: "abc1"}
	data, _ := json.Marshal(url)
	res.Write(data)
}
