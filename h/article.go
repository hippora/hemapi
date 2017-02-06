package h

import (
	"github.com/kataras/iris"
	"github.com/hippora/hemapi/m"
	"log"
)

func GetArticle(ctx *iris.Context)  {
	articles,err := m.GetArticles()
	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON(iris.StatusOK,articles)
}