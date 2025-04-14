package main

import (
	"smsforwarder-manager/models"
	"smsforwarder-manager/router"
)

func main() {
	// init db
	models.NewDB()

	r := router.App()

	r.Run(":801")
}
