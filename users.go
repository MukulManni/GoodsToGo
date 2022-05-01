package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Fname   string
	Lname   string
	Email   string
	Mobile  string
	Stateid string
	Pass    string
}

var userList = []User{{"admin", "admin", "admin@goodstogo.com", "9876543210", "1", "admin@123"}}

const (
	username = "user"
	password = "pass"
)

func createUser(c *gin.Context) {
	fname := c.PostForm("FirstName")
	lname := c.PostForm("LastName")
	email := c.PostForm("EmailId")
	mobile := c.PostForm("MobileNo")
	stateid := c.PostForm("StateId")
	pass := c.PostForm("password")
	repass := c.PostForm("repass")

	if pass != repass {
		c.HTML(
			http.StatusOK,
			"SignUp.html",
			gin.H{
				"error": "Password not same",
			},
		)
		return
	}

	for _, user := range userList {
		if user.Email == email {
			c.HTML(
				http.StatusOK,
				"SignUp.html",
				gin.H{
					"error": "Email already registered",
				},
			)
			return
		}
	}

	user := User{fname, lname, email, mobile, stateid, pass}

	userList = append(userList, user)

	c.Redirect(http.StatusFound, "/login")
}
