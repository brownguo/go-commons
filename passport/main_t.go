package main

import (
	"fmt"
	"github.com/kataras/iris"
)

func main(){

	fmt.Printf("Hello Test Page!")

	app := iris.New()

	app.RegisterView(iris.HTML("./views",".html"))

	app.Get("/",func(ctx iris.Context){
		ctx.ViewData("messages","MainTestPage!,Based var!")
		ctx.View("main_test.html")
	})

	app.Run(iris.Addr(":8080"))

}