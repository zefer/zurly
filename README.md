# Zurly

A very simple URL shortening service.

Go + Redis.

Built to help me learn Go, so expect scaryness.

Deliberately using net/http rather than a web framework, in an attempt to learn.

# Short Urls

A unique ID is generated for each short URL. This ID is the hex value of the
total number of urls stored in Redis. This total is implemented as a simple
counter.

Path /:ID will redirect to the corresponding long url for the given ID.

# API

I intend to implement both a JSON and a simple HTTP get/url based API using
content negotiation.
