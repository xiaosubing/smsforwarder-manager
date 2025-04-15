package router

import (
	"github.com/gin-gonic/gin"
	"smsforwarder-manager/middleware"
	"smsforwarder-manager/service"
)

func App() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())

	r.POST("/api/login", service.Login)
	r.POST("/sms/query", service.GetMessages)
	r.POST("/api/sendMessages", service.SendMessages)
	r.POST("/api/getPhones", service.GetPhones)
	r.POST("/api/forwarderMessage", service.ForwarderMessage)
	r.POST("/api/getCode", service.GetVerifyCode)
	return r
}
