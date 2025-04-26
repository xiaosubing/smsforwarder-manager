package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"smsforwarder-manager/models"
)

func ChangePwd(c *gin.Context) {
	changePwdData := new(ChangePassword)
	err := c.ShouldBind(changePwdData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  400,
		})
		fmt.Println(err)
		return
	}
	fmt.Println(changePwdData)
	err = models.ChagePwd(changePwdData.Username, changePwdData.OldPassword, changePwdData.NewPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
	fmt.Println("修改成功")
}
