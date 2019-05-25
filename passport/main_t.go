package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"go-commons/passport/controller"
)

func main(){

	//fmt.Printf("Hello Test Page!")

	//create iris object!

	app := iris.New()

	//view path.
	app.RegisterView(iris.HTML("./views",".html"))

	//not found page.
	app.OnErrorCode(iris.StatusNotFound, notFound)

	//webRoot
	app.Get("/",func(ctx iris.Context){
		ctx.ViewData("messages","MainTestPage!,Based var!")
		ctx.View("main_test.html")
	})

	//return json.
	app.Get("/getjson",func(ctx iris.Context){
		tmp := "jsonstring."
		ctx.JSON(tmp)
	})

	//controller
	app.Get("/c",controller.Authentication)

	//a get request.
	app.Get("/getRequest", func(context context.Context) {
		//获取请求路径
		path := context.Path()
		//输出请求日志
		app.Logger().Info(path)
		context.WriteString("请求路径：" + path)
	})

	//a post request.
	app.Handle("POST", "/user/info", func(context context.Context) {
		context.WriteString(" User Info is Post Request , Deal is in handle func \n")
	})

	app.Run(iris.Addr(":8080"))


	/**
		返回string类型数据
		context.WriteString("hello world")
		返回json格式的数据
		context.JSON(iris.Map{"message": "hello word", "requestCode": 200})
		返回xml格式的数据
		context.XML(Person{Name: "Davie", Age: 18})
		返回html格式数据
		context.HTML("<h1> Davie, 12 </h1>")
	 */

}

func notFound(ctx iris.Context) {
	// 出现 404 的时候，就跳转到 $views_dir/errors/404.html 模板
	ctx.View("notfound.html")
}