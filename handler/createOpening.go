package handler

import "github.com/gin-gonic/gin"

func CreateOpeningHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "POST Opening",
	})
}
