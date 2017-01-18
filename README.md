# short-url
Simple short url service

# TODO

* [ ] auto redirect for original url

# Pitfalls

* slice in Golang will auto grow length
* atomic inc should be used carefully
* error handling
* must wrap a new ResponseWriter to get the status code