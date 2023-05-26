package main

import (
	// "fmt"
	"fmt"
	"newsfeeder/httpd/handler"
	"newsfeeder/repository/newsfeeds"

	"github.com/gin-gonic/gin"
)

func main(){
	feed := newsfeeds.New()
	// feed.AddItem(
	// 	newsfeeds.Item{
	// 		Title: "hello",
	// 		Body: "World",
	// 	},
	// );
	r:=gin.Default();  
    fmt.Print(feed.GETAll())
	r.GET("/",handler.PingGet())
    r.GET("/news",handler.GetNews(feed))
	r.POST("/addNews",handler.PostNews(feed))
	r.Run(":8080")
}