package main

import (
	"newsfeeder/httpd/handler"
	"github.com/gin-gonic/gin"
)



func main(){
	r:=gin.Default();  
    
	r.GET("/",handler.PingGet)

	r.Run(":8080")

}