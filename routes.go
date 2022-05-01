package main

func loadTemplates() {

	r.LoadHTMLFiles("static/home/index.html", "static/login/login.html", "static/login/Login/SignUp.html")

}

func intitalizeRoutes() {

	r.Static("/static/home", "static/home")
	r.Static("/static/login", "static/login")

	r.GET("/", homePage)

	r.GET("/login", login)
	r.POST("/login", validateUser)

	r.GET("/register", signup)
	r.POST("/register", createUser)
}
