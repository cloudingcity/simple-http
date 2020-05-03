# Simple HTTP

Simple HTTP server to be used to test load balancer.

## Usage

```shell script
$ docker run -d -p 8080:80 ghost0436/simple-http
$ curl :8080
{"hostname":"6c32cb114ebd"} # return the current container hostname
```
