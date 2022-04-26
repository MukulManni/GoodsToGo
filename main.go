package main

import "github.com/gin-gonic/gin"

var r *gin.Engine

func main() {
	r = gin.Default()

	loadTemplates()
	intitalizeRoutes()

	r.Run(":80")
}
