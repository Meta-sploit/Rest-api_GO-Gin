package handler

import (
	"net/http"
	"newsfeeder/repository/newsfeeds"

	"github.com/gin-gonic/gin"
)

func GetNews() gin.HandlerFunc{
	feed := newsfeeds.New()
	result := feed.GETAll()
	return func (c *gin.Context)  {
		c.JSON(http.StatusOK,result)
	}
}