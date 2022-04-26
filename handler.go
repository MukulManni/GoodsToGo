package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func homePage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		nil,
	)
}

func login(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"login.html",
		nil,
	)
}

func signup(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"SignUp.html",
		nil,
	)
}
