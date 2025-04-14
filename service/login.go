package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"smsforwarder-manager/middleware"
	"smsforwarder-manager/models"
)

func Login(c *gin.Context) {
	recInfos := new(LoginRequest)
	err := c.ShouldBind(recInfos)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  400,
		})
		fmt.Println(err)
		return
	}

	user := models.GetUserInfo(recInfos.Username, recInfos.Password)

	if user.Username == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "404",
		})
		return
	}

	if user.Username != recInfos.Username || user.Password != recInfos.Password {
		c.JSON(http.StatusNotFound, gin.H{
			"code": -1,
			"msg":  "404",
		})
		return
	}

	fmt.Println("验证通过")
	// 生成token
	token, err := middleware.GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": -1,
			"msg":  "404",
		})
		return
	}
	refreshToken, err := middleware.GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": -1,
			"msg":  "404",
		})
		return
	}

	date := &LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "登录成功",
		"data": *date,
	})
}
