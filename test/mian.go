package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"test/model"
)


var db *sql.DB

func main(){
	r:=gin.Default()
    r.POST("/register",Registe)
    r.POST("/login",Login)
	r.POST("/sendmasge",SendMsg)
	r.GET("/getmasge",GetMsg)
	r.Run(":8080")
}



func Registe(c *gin.Context){
username:=c.PostForm("username")
password:=c.PostForm("password")

fmt.Println("user:"+username+password)
if model.UserSignup(username,password){
c.JSON(500,gin.H{"status":http.StatusInternalServerError,"message":"数据库Insert报错"})
}else {
c.JSON(200, gin.H{"status": http.StatusOK, "message": "注册成功"})
}
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if model.UserSignin(username, password) {
		c.SetCookie("username", username, 10, "localhost:8080", "localhost", false, true)
		c.JSON(200, gin.H{"status": http.StatusOK, "message": "登录成功"})
	} else {
		c.JSON(403, gin.H{"status": http.StatusForbidden, "message": "登录失败，用户名或密码错误"})
	}
}


func SendMsg(c *gin.Context){
	username,_:=c.Cookie("username")
	fmt.Println("username"+username)
	/*if err != nil{
		c.JSON(500,gin.H{"status": http.StatusForbidden,"message":"cookie读取失败"})
		return
	}*/
	message:=c.PostForm("message")
	if model.SendMessage(username,message){
		c.JSON(200, gin.H{"内容":message,"用户名":username})
	}else {
		c.JSON(403, gin.H{"status": http.StatusForbidden, "message": "发送失败"})
	}
}

func GetMsg(c *gin.Context){

	err,msg:=model.GetMesage(10)
	if err !=nil{
		c.JSON(500,gin.H{"status": http.StatusInternalServerError,"message":"数据库读取失败"})
	}
	c.JSON(200,gin.H{"Data":msg})
}