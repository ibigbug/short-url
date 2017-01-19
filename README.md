# short-url

Simple short url service

# Start the service

```
$ git clone https://github.com/ibigbug/short-url.git
$ cd short-url
$ sh ./scripts/start.sh
```

# Usage

```
$ curl -sX POST -H 'Content-Type: application/json' 'localhost:8000/shorten' -d '{"url":"http://a.very.long.url"}'
{"short": "http://localhost/1"}
```

```
$ curl -sX GET -H 'Content-Type: application/json' 'localhost:8000/original' -d '{"short":"http://localhost/1"}'
{"original": "http://a.very.long.url"}
```

# TODO

* [ ] auto redirect for original url
* [ ] replace in-memory storage with Redis or other persistent stroage
* [ ] implement querystring support

# Pitfalls

* slice in Golang will auto grow length
* atomic inc should be used carefully
* must wrap a new ResponseWriter to get the status code
