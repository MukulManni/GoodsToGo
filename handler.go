package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func homePage(c *gin.Context) {

	session := sessions.Default(c)

	user := session.Get("username")

	if user == nil {
		c.HTML(
			http.StatusOK,
			"index.html",
			nil,
		)
	} else {

		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"username": user,
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

	for _, u := range userList {
		if u.Email == user {
			if u.Pass == pass {

				session := sessions.Default(c)

				session.Set(username, user)

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
