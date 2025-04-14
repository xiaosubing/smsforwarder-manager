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
	r.POST("/api/getMessages", service.GetMessages)
	r.POST("/api/getPhones", service.GetPhones)
	r.POST("/api/sendMessages", service.SendMessages)
	r.POST("/api/forwarderMessage", service.ForwarderMessage)
	return r
}
