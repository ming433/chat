package controller

import (
	"chat/apps/api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong1",
	})
}

func GetUserList(c *gin.Context) {
	userList := models.GetUserList()
	c.JSON(http.StatusOK, gin.H{
		"data": userList,
	})
}
