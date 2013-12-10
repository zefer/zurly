package main

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
)

type Url struct {
	Id      string
	LongUrl string
}

type UrlInput struct {
	Url string
}

func findUrl(id string) (*Url, *Error) {
	log.Printf("Finding url: %v", id)
	str, _ := redis.Get("zurl:" + id)
	log.Printf("Found url: %v", str)
	if str == nil {
		return nil, &Error{Message: "Couldn't find this URL :("}
	} else {
		url := &Url{}
		json.Unmarshal([]byte(str), &url)
		return url, nil
	}
}

func (this *Url) json() []byte {
	data, _ := json.Marshal(this)
	return data
}

func (this *UrlInput) validate() (bool, *Error) {
	match, _ := regexp.MatchString(`^(https?:\/\/)([\da-z\.-]+)\.([a-z\.]{2,6})([\?#=&\/\w \.\-]*)*\/?$`, this.Url)
	if match {
		return true, &Error{Message: ""}
	} else {
		return false, &Error{Message: "not a valid url"}
	}
}

func (this *Url) save() {
	this.generateId()
	log.Printf("Saving url: %v", this.Id)
	redis.Set("zurl:"+this.Id, []byte(this.json()))
}

// set the ID to the hex value of the counter
func (this *Url) generateId() {
	count, _ := redis.Incr("zurl:counter")
	this.Id = fmt.Sprintf("%x", count)
}
