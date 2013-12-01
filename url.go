package main

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
)

type Zurl struct {
	Id      string
	LongUrl string
}

func findZurl(id string) (*Zurl, *Error) {
	log.Printf("Finding zurl: %v", id)
	str, _ := redis.Get("zurl:" + id)
	log.Printf("Found zurl: %v", str)
	if str == nil {
		return nil, &Error{Message: "Couldn't find this URL :("}
	} else {
		zurl := &Zurl{}
		json.Unmarshal([]byte(str), &zurl)
		return zurl, nil
	}
}

func (this *Zurl) json() []byte {
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
	redis.Set("zurl:"+this.Id, []byte(this.json()))
}

// set the ID to the hex value of the counter
func (this *Zurl) generateId() {
	count, _ := redis.Incr("zurl:counter")
	this.Id = fmt.Sprintf("%x", count)
}
