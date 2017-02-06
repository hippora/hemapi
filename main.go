package main

import (
	"github.com/kataras/iris"
	"github.com/hippora/hemapi/h"
)


func init() {
	iris.Get("/a/article", h.GetArticle)
	iris.Post("/a/user", h.EnrolUser)
}

func main() {
	//iris.OptionGzip(true)
	iris.Listen(":8080")
}
