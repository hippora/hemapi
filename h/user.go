package h

import "github.com/kataras/iris"

func EnrolUser(ctx *iris.Context)  {
	ctx.SetCookieKV("name","hippo")
	ctx.JSON(iris.StatusOK,map[string]string{"register": "user"})
}
