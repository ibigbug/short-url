# short-url

[![Build Status](https://travis-ci.org/ibigbug/short-url.svg?branch=master)](https://travis-ci.org/ibigbug/short-url)
[![Coverage Status](https://coveralls.io/repos/github/ibigbug/short-url/badge.svg?branch=master)](https://coveralls.io/github/ibigbug/short-url?branch=master)

Simple short url service

# Start the service

## Docker

```
$ docker run -e PORT=80 -e ADDR=0.0.0.0 -p 8000:80 -e SITE_URL=YOUR_DOMAIN --rm -it ibigbug/short-url
```

## Clean Virtual Machine

```
$ git clone https://github.com/ibigbug/short-url.git
$ cd short-url
$ sh ./scripts/bootstrap.sh
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
