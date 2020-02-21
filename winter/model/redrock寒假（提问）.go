package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ask(c *gin.Context){
	username,err:=c.Cookie("username")
	ask:=c.PostForm("ask")
	if err!=nil {
		fmt.Println("cookie读取失败")
	}
	if Question(username,ask) {
		question_number:=Select_question_number(username)
		Update_q_number(username,question_number)
		c.JSON(200,gin.H{"status":http.StatusOK,"message":"成功发表提问"})
	}else {
		c.JSON(200, gin.H{"status": http.StatusOK, "message": "数据库insert报错"})
	}
}

func Question(username string,ask string)bool{
	stmt,err:=DBConn().Prepare("insert into question(username,ask) values(?,?)")
	if err!=nil {
		fmt.Println("insert提问时出现问题")
		return false
	}
	defer stmt.Close()
	stmt.Exec(username,ask)
	return true
}

func Select_question_number(username string)(int){
	stmt,err:=db.Query("select question_number from user where username=?",username)
	if err!=nil {
		fmt.Println("查找提问数时失败")
	}
	defer stmt.Close()
	var question_number int
	stmt.Scan(&question_number)
	return question_number
}

func Update_q_number(username string,question_number int){
	question_number+=1
	stmt,err:=db.Prepare("update user set question_number='?'value(?)where username='?'value(?)")
	if err!=nil {
		fmt.Println("更新失败")
	}
	defer stmt.Close()
	stmt.Exec(question_number,username)
}