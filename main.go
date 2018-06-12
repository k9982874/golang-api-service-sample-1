package main

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/k9982874/golang-api-service-sample-1/controllers"
)

func main() {
	r := gin.Default()

	controllers.NewApplicationController(r)

	r.Run(":9999")
}
