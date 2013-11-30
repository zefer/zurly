package main

import (
	"log"
	"net/http"
	"encoding/json"
	"fmt"
	"regexp"
)

type Zurl struct {
	Id string
	LongUrl string
}
func (zurl Zurl) validate() (bool, *Error) {
	match, _ := regexp.MatchString(`^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$`, zurl.LongUrl)
	if match {
		return true, &Error{Message: ""}
	} else {
		return false, &Error{Message: "not a valid url"}
	}
}

type Error struct {
	Message string
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
	zurl := &Zurl{Id: "abc1", LongUrl: req.FormValue("url")}
	valid, err := zurl.validate()
	if valid {
		data, _ := json.Marshal(zurl)
		res.WriteHeader(201)
		res.Write(data)
	} else {
		data, _ := json.Marshal(err)
		res.WriteHeader(422)
		res.Write(data)
	}
}
