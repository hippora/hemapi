package c

import "github.com/kataras/iris"

func api_register_user(ctx *iris.Context)  {
	ctx.SetCookieKV("name","hippo")
	ctx.JSON(iris.StatusOK,map[string]string{"register": "user"})
}

func api_article(ctx *iris.Context)  {
	ctx.JSON(iris.StatusOK,map[string]string{"api": "article"})
}