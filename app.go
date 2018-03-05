package main

import (
	"github.com/kataras/iris"
)

// Serve using a host:port form.
var addr = iris.Addr("localhost:8080")

func main() {
	app := iris.New()

	// GET: http://localhost:8080
	app.Get("/", index)

	// GET: http://localhost:808/print
	app.Get("/print", iris.Gzip, article)

	// Now listening on: http://localhost:8080
	// Application started. Press CTRL+C to shut down.
	app.Run(addr)
}

func print(ctx iris.Context) {
  ctx.JSON(iris.Map{"message": "Print"})
}

func index(ctx iris.Context) {
	ctx.JSON(iris.Map{"message": "Hello World"})
}
