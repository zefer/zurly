package main

import (
	"github.com/codegangsta/martini"
)

func main() {
	m := martini.New()
	r := martini.NewRouter()
	m.Use(martini.Logger())
	m.Use(martini.Recovery())
	m.Action(r.Handle)

	r.Get("/", Welcome)
	r.Get("/:id", ExpandUrl)
	m.Run()
}

func Welcome() (int, string) {
	return 200, "Hello! Zurl is a URL shortener service."
}

func ExpandUrl(params martini.Params) (int, string) {
	return 200, "ID: " + params["id"]
}
