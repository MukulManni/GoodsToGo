package main

import (
	"os"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func main() {
	r = gin.Default()

	port := os.Getenv("PORT")

	r.Use(sessions.Sessions("usersession", sessions.NewCookieStore([]byte("secret"))))

	loadTemplates()
	intitalizeRoutes()

	r.Run(":" + port)
}
