package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"test/jwt"
	"test/resps"
)


var(
	DB *gorm.DB
)

func init(){
	mysql, err := gorm.Open("mysql", "root:lhr20001221@(127.0.0.1:3306)/firstweek?charset=utf8&parseTime=true")
	if err!=nil{
		log.Fatal(err.Error())
	}
	DB=mysql
}

type LoginForm struct {
	username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type user struct {
	gorm.Model
	Username string
	Password string
}

func Login(c *gin.Context){
	var l LoginForm
	l.username=c.PostForm("username")
	l.Password=c.PostForm("password")
	DB.Where("")
	if err:=c.ShouldBindJSON(&l);err!=nil{
		resps.ErrorWeb(c)
		return
	}
	var u user
	DB.Where(user{
		Username:l.username,
		Password:l.Password,
	}).First(&u)
	if u.ID==0 {
		resps.Error(c,1003,"can not find this one")
	}else {
		resps.OkData(c,u)
	}
	jwt.Create_jwt(l.username)
}
