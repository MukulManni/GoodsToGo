package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func main() {
	r = gin.Default()

	r.Use(sessions.Sessions("usersession", sessions.NewCookieStore([]byte("secret"))))

	loadTemplates()
	intitalizeRoutes()

	r.Run(":80")
}
