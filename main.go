package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"aatrox/chess/controllers"
)

func main() {
	var app = iris.Default()

	mvc.Configure(app.Party("/games"), func(app *mvc.Application) {
		app.Handle(new(controllers.GameController))
	})

	app.Listen(":8080")
}
