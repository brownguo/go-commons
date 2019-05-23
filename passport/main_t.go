package main

import (
	"fmt"
	"github.com/kataras/iris"
	"go-commons/passport/controller"
)

func main(){

	fmt.Printf("Hello Test Page!")

	app := iris.New()

	app.RegisterView(iris.HTML("./views",".html"))

	app.Get("/",func(ctx iris.Context){
		ctx.ViewData("messages","MainTestPage!,Based var!")
		ctx.View("main_test.html")
	})

	app.Get("/getjson",func(ctx iris.Context){
		tmp := "jsonstring."

		ctx.JSON(tmp)
	})

	app.Get("/c",controller.Authentication)

	app.Run(iris.Addr(":8080"))

}