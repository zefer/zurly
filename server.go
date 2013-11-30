package main

import (
	"log"
	"net/http"
	"encoding/json"
	"fmt"
	"regexp"
	"github.com/hoisie/redis"
)

var redisClient redis.Client

type Zurl struct {
	Id string
	LongUrl string
}

func (this *Zurl) json() ([]byte) {
	data, _ := json.Marshal(this)
	return data
}

func (this *Zurl) validate() (bool, *Error) {
	match, _ := regexp.MatchString(`^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$`, this.LongUrl)
	if match {
		return true, &Error{Message: ""}
	} else {
		return false, &Error{Message: "not a valid url"}
	}
}

func (this *Zurl) save() {
	this.generateId()
	log.Printf("Saving zurl: %v", this.Id)
	redisClient.Set("zurl:" + this.Id, []byte(this.json()))
}

// set the ID to the hex value of the counter
func (this *Zurl) generateId() {
	count, _ := redisClient.Incr("zurl:counter")
	this.Id = fmt.Sprintf("%x", count)
}

type Error struct {
	Message string
}
func (this *Error) json() ([]byte) {
	data, _ := json.Marshal(this)
	return data
}

func initCounter() {
	counter, _ := redisClient.Get("zurl:counter")
	if counter == nil {
		log.Print("Initialising counter")
		redisClient.Set("zurl:counter", []byte("0"))
	} else {
		log.Printf("Counter is " + string(counter))
	}
}

func main() {
	initCounter()
	http.HandleFunc("/", Root)

	log.Print("Listening")
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
	zurl := &Zurl{Id: "", LongUrl: req.FormValue("url")}
	valid, err := zurl.validate()
	if valid {
		zurl.save()
		res.WriteHeader(201)
		res.Write(zurl.json())
	} else {
		res.WriteHeader(422)
		res.Write(err.json())
	}
}
