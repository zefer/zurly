package main

import (
	"os"
	"log"
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/hoisie/redis"
	"strings"
)

var redisClient redis.Client

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
	address := os.Getenv("REDIS_ADDRESS")
	if address != "" {
		log.Printf("Connecting to redis at %v", address)
		redisClient = redis.Client{
			Addr:     address,
			Password: os.Getenv("REDIS_PASSWORD"),
		}
	}

	initCounter()
	http.HandleFunc("/", Root)
	port := os.Getenv("PORT")
	log.Print("Listening on port " + port)
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Fatal("Error: %v", err)
	}
}

func Root(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		if req.URL.Path == "/" {
			Welcome(res, req)
		} else {
			Expand(res, req)
		}
	case "POST":
		if req.URL.Path != "/" {
			http.NotFound(res, req)
			return
		}
		Shorten(res, req)
	default:
		res.WriteHeader(405)
	}
}

func Welcome(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(res, "Hello! Zurl is a URL shortener service.")
}

func Expand(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	id := strings.Replace(req.URL.Path, "/", "", -1)
	zurl, err := findZurl(id)
	if err == nil {
		http.Redirect(res, req, zurl.LongUrl, 302)
	} else {
		res.WriteHeader(404)
		res.Write([]byte(err.Message))
	}
}

func Shorten(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	zurl := &Zurl{LongUrl: req.FormValue("url")}
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
