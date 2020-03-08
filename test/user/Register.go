package user

import (
	"github.com/gin-gonic/gin"
	"test/resps"
)

func Register(c *gin.Context){
	var l  LoginForm
	l.username=c.PostForm("username")
	l.Password=c.PostForm("password")
	if err:=c.ShouldBindJSON(&l);err!=nil{
		resps.ErrorWeb(c)
	}
	var u user
	DB.Where("username=?",l.username).First(&u)
	if u.ID==0 {
		resps.Error(c,1003,"fail to do this")
	}
	u = user{
		Username: l.username,
		Password: l.Password,
	}
	DB.Create(&u)
	resps.OkData(c,gin.H{"message":"quite ok"})
}
