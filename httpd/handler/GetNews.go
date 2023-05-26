package handler

import (

	"net/http"
	"newsfeeder/repository/newsfeeds"
	"github.com/gin-gonic/gin"
)

func GetNews(feed *newsfeeds.NewsFeeds) gin.HandlerFunc{
	return func (c *gin.Context)  {
		result := feed.GETAll()
		c.JSON(http.StatusOK,result)
	}
}