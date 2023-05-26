package handler

import (
	"fmt"
	"net/http"
	"newsfeeder/repository/newsfeeds"

	"github.com/gin-gonic/gin"
)

type NewsFeedsPost struct{
	Title string `json:"title"`
	Body string `json:"body"`
}


func PostNews(feed *newsfeeds.NewsFeeds) gin.HandlerFunc{
	return func (c *gin.Context)  {
        responseBody :=  NewsFeedsPost{}
		c.Bind(&responseBody)
    
		news := newsfeeds.Item{
			Title: responseBody.Title,
			Body: responseBody.Body,
		}

		feed.AddItem(news)
        fmt.Print(feed)
		c.Status(http.StatusNoContent)
	}
}