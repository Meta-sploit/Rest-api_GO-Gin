package main

import (
	"fmt"
	"newsfeeder/repository/newsfeeds"
)

// "newsfeeder/httpd/handler"
// "github.com/gin-gonic/gin"



func main(){
	// r:=gin.Default();  
    
	// r.GET("/",handler.PingGet)

	// r.Run(":8080")
	feed:= newsfeeds.New()
	fmt.Print(feed)

	feed.AddItem(newsfeeds.Item{"hello","world"})
	fmt.Print(feed.GETAll())

}