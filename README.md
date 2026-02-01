# HttpPlusPlus

A dead simple ```net/http``` wrapper.

**Note:** The middleware api will change soon.

## Installation
```sh
go get github.com/chuckcrump/httplusplus
```

## Features
- Context w/support for text, json, and error responses
- Set and get key-value pairs in context
- Middleware
- Nested routers 

## Basic Example
```go
package main

import (
	"net/http"

	hpp "github.com/chuckcrump/httplusplus"
)

func HelloMiddleware(next hpp.Handler) hpp.Handler {
  return func (ctx *hpp.Context) error {
    fmt.Println("A route is being called...")
    next(ctx)
    return nil
  }
}

func Hello(ctx *hpp.Context) error {
  name := ctx.Request.PathValue("name")
  return ctx.Text(http.StatusOK, "Hallo, "+name+"!")
}

func main() {
	app := hpp.NewRouter()

	app.Route("GET /hello/{name}", hpp.Use(Hello, HelloMiddleware //any other middlewares))
	//NOTE in the future it will look like
	//app.
	//  Route("GET /hello/{name}", Hello).
	//  Middlewares(HelloMiddleware //any other middlewares)
  
	hpp.StartApp("127.0.0.1:3001", app)
}
```
