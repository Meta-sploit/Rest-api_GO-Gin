package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingGet(c *gin.Context){
     c.JSON(http.StatusOK,"hello world")
}