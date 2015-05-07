# go-wave-workshop

Something that can be `go get`'d to bootstrap our Wave workshop

## Branches

* [`master`](https://github.com/NickPresta/go-wave-workshop/tree/master) is our first attempt at learning Go
* [`engineered`](https://github.com/NickPresta/go-wave-workshop/tree/engineered) is our cleaned up attempt with some basic tests and structure/packages

## How to run

1. Start the server: `go run main.go`
1. Make a POST request to the server:

```
curl -X POST -d '{"amount": "100", "from": "USD", "to": "CAD"}' http://localhost:12345/convert
```
