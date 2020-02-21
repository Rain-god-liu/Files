package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)



type Informations struct {
	information string
	id int
}

func Search(c *gin.Context){
	var informations []Informations
	word:=c.PostForm("keyword")
	stmt,err:=db.Query("select infomation,id frome search where keyword like '%word%=?'",word)
	if err!=nil {
		log.Fatal("查询时出现错误")
	}
	defer stmt.Close()
	for stmt.Next() {
		var information Informations
		stmt.Scan(&information.information,&information.id)
		if err!=nil{
			fmt.Println("fail to do this")
		}
		informations=append(informations,information)
	}
	if err!=nil {
		c.JSON(500,gin.H{"status": http.StatusInternalServerError,"message":"数据库读取失败"})
	}
	c.JSON(200,gin.H{"Data":informations})
}

func Insert_information(keyword string,information string)bool{
	stmt,err:=DBConn().Prepare("insert into search(keyword,information) values (?,?)")
	if err!=nil {
		log.Fatal("添加信息时出现错误")
		return false
	}
	defer stmt.Close()
	_,err=stmt.Exec(keyword,information)
	if err!=nil {
		fmt.Println("添加信息时发生错误")
		return false
	}
    return true
}
