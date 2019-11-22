package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var USER map[string]User = make(map[string]User)

var user User

func main() {

	origin:= gin.Default()

	origin.POST("/register", func(context *gin.Context) {
		err := context.Bind(&user)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}

		if _, ok := USER[user.username]; ok {
			context.JSON(200, gin.H{
				"message": "用户名" + user.username + "已存在",
			})
		} else {
			USER[user.username] = user
			context.JSON(200, gin.H{
				"message": "注册成功",
			})
		}

	})
	origin.POST("/login", func(context *gin.Context) {
		err := context.Bind(&user)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}

		if a, ok := USER[user.username]; ok && a .password == user.password {
			context.JSON(http.StatusOK, gin.H{
				"username": a .username,
				"password": a .password,
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"message": "账号或者密码有误",
			})
		}
	})
	origin.Run(":8080")
}

type User struct {
	username string `form:"name"`
	password string `form:"password"`
}

