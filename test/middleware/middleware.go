package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"test/jwt"
	"test/resps"
)

func User(c *gin.Context) {
	token:= c.GetHeader("Authorization")
	fmt.Println(token)
	if len(token)<7 {
		resps.Error(c, 1004, "token error")
		c.Abort()
		return
	}
	token2 := token[7:]
	username, err := jwt.Makesure_token(token2)
	fmt.Println(err)
	if err != nil {
		resps.Error(c, 1005, "token error")
		c.Abort()
		return
	}
	c.Set("username", username)
	c.Next()
	fmt.Println(token2)
	return
}