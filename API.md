FORMAT: X-1A

# Zurly URL shortener service
Zurly is a simple URL shorter service, built to help me learn Go.

# Group URLs
URL resources.

## URL [/{id}]
A URL resource.

+ Parameters
  + id (string, `a1`) ... The id of the URL.

+ Model (application/json)

    ```js
    {
      "Id": "1",
      "LongUrl": "http://original.url"
    }
    ```

### Expand a short URL [GET]
Redirect to the original (long) URL.

+ Request redirect to the long URL
  + Headers
      Accept: text/html

+ Response 302 (text/html)
  + Headers
      Location: http://original.url
  + Body
      <a href="http://original.url">Found</a>.

### Expand a short URL [GET]
Return the URL resource as JSON

+ Request JSON Message
  + Headers
      Accept: application/json

+ Response 200

  [URL][]

## URL [/]

### Shorten a URL [POST]
Shorten a URL with a JSON POST request

+ Request
  + Headers
      Accept: application/json
  + Body

    ```js
    {
      "url": "http://original.url"
    }
    ```

+ Response 201

  [URL][]
