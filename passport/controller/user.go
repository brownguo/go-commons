package controller

import (
	"github.com/kataras/iris"
)


//type UserController struct {
//
//}

//func (m *UserController) Authentication(b mvc.BeforeActivation){
//	fmt.Println("Hello Controller!")
//}

func Authentication(ctx iris.Context) {
	//user_id := ctx.Values().Get("auth_user_id").(uint)
	//ctx.JSON(user_id)
	ctx.JSON("get controller success!")
}