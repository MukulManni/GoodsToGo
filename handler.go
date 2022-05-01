package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	username = "user"
	password = "pass"
)

func homePage(c *gin.Context) {

	session := sessions.Default(c)

	user := session.Get(username)
	fmt.Println("home:", user)

	if user == nil {
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"check": "Sign Up",
			},
		)
	} else {

		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"check": user,
			},
		)
	}

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

func validateUser(c *gin.Context) {
	user := c.PostForm("username")
	pass := c.PostForm("password")

	fmt.Println(user, pass)

	for _, u := range userList {
		if u.Email == user {
			if u.Pass == pass {

				session := sessions.Default(c)

				session.Set(username, u.Fname)
				session.Save()

				fmt.Println(session.Get(username))

				c.Redirect(http.StatusFound, "/")

				return
			} else {
				c.HTML(
					http.StatusOK,
					"login.html",
					gin.H{
						"error": "Incorrect Password.",
					},
				)

				return
			}
		}
	}

	c.HTML(
		http.StatusOK,
		"login.html",
		gin.H{
			"error": "Email not found.",
		},
	)
}
