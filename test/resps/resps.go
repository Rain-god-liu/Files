package resps

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func OkWeb(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{"code":1000,"message":"ok"})
}

func OkData(c *gin.Context,data interface{}){
	c.JSON(http.StatusOK,gin.H{"code":1001,"message":"ok","Data":data})
}

func ErrorWeb(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{"code":1002,"message":"some web error happen"})
}

func Error(c *gin.Context,code int,msg string){
    c.JSON(http.StatusOK,gin.H{"code":code,"message":msg})
}