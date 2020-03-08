package contoller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"test/middleware"
	"test/user"
)


func Set_user_router(r *gin.Engine){
	r.POST("/register",user.Register)
	r.POST("/login",user.Login)
	r.GET("/check", middleware.User,func(c *gin.Context) {
		fmt.Println("token is quite ok")
	})
}