package main

import "github.com/kataras/iris"



func init() {
	iris.Get("/a/article",api_article)
	iris.Get("/a/user",api_register_user)

}


func main() {
	iris.Config.Gzip = true
	iris.Listen(":8080")
}
