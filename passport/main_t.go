package main

import (
	"github.com/kataras/iris"
	"go-commons/passport/controller"
)

func main(){

	//fmt.Printf("Hello Test Page!")

	app := iris.New()

	app.RegisterView(iris.HTML("./views",".html"))
	app.OnErrorCode(iris.StatusNotFound, notFound)

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

func notFound(ctx iris.Context) {
	// 出现 404 的时候，就跳转到 $views_dir/errors/404.html 模板
	ctx.View("notfound.html")
}