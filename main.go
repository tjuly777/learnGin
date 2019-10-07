package main

import (
	"./routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")
	router.Static("/assets", "./assets")

	//ルーティング
	user := router.Group("/user")
	{
		user.POST("/signup", routes.UserSignUp)
		user.POST("/login", routes.UserLogin)
	}

	router.GET("/", routes.Home)
	router.GET("/login", routes.Login)
	router.GET("/signup", routes.SignUp)
	router.NoRoute(routes.NoRoute)

	//ポートを指定してサーバー起動
	router.Run(":5000")
}
