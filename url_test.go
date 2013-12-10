package main

import (
	"testing"
)

func TestUrlInput(t *testing.T) {
	var validUrls = []string{
		"http://zefer.co",
		"http://zefer.co/",
		"http://zefer.co/banana-jungle-1",
		"http://zefer.co/gorilla?banana=false&jungle=massive",
		"http://zefer.co?gorilla=true#bananas",
	}
	var invalidUrls = []string{
		"http://zefer.c",
		"http://zefer/",
		"zefer.co",
		"httx://zefer.co",
		"http://.co",
	}

	for i := range validUrls {
		u := UrlInput{Url: validUrls[i]}
		valid, _ := u.validate()
		if valid == false {
			t.Errorf("Url %v incorrectly identified as invalid", validUrls[i])
		}
	}

	for i := range invalidUrls {
		u := UrlInput{Url: invalidUrls[i]}
		valid, _ := u.validate()
		if valid == true {
			t.Errorf("Url %v incorrectly identified as valid", invalidUrls[i])
		}
	}
}
