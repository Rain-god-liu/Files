package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"test/contoller"
)

func main(){
	r:=gin.Default()
    contoller.Set_user_router(r)
    r.Run(":8080")
}
