package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Write(c *gin.Context){
	username,err:=c.Cookie("username")
	if err!=nil {
		fmt.Println("获取cookie时出现错误")
	}
	article:=c.PostForm("article")
	if insert_article(username,article) {
		article_number:=Select_article_number(username)
		Update_a_number(username,article_number)
		c.JSON(200,gin.H{"status":http.StatusOK,"message":"发表文章成功"})
	}else {
		c.JSON(500,gin.H{"status":http.StatusInternalServerError,"message":"数据库insert报错"})
	}
}

func insert_article(username string,article string)bool{
	stmt,err:=DBConn().Prepare("insert into article(username,article) values(?,?)")
	if err!=nil {
		fmt.Println("insert报错")
		return false
	}
	defer stmt.Close()
	stmt.Exec(username,article)
	return true
}

func Select_article_number(username string)(int){
	stmt,err:=db.Query("select question_number from user where username=?",username)
	if err!=nil {
		fmt.Println("查找提问数时失败")
	}
	defer stmt.Close()
	var article_number int
	stmt.Scan(&article_number)
	return article_number
}

func Update_a_number(username string,article_number int){
	article_number+=1
	stmt,err:=db.Prepare("update user set article_number='?'value(?)where username='?'value(?)")
	if err!=nil {
		fmt.Println("更新失败")
	}
	defer stmt.Close()
	stmt.Exec(article_number,username)
}