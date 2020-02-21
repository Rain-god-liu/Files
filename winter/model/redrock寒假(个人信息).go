package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Userinfor struct {
	username string
	question_number int
	focus_number int
}

type name struct {
	username string
	id int
}

type article struct {
	username string
	atricle string
}

type ask struct {
	username string
	ask string
}

func User_information(c *gin.Context){
	var err error
	var userinfor Userinfor
	username,err1:=c.Cookie("username")
	if err1!=nil {
		fmt.Println("获取cookie的username时出现错误")
	}
	//获取用户的username，提问数，关注人数信息
	err,userinfor=Select_userinfor(username)
	if err!=nil {
		c.JSON(500,gin.H{"status": http.StatusInternalServerError,"message":"数据库读取失败"})
	}
	c.JSON(200,gin.H{"Data":userinfor})
	//获取关注者的名字和id
	err2,focusman_name:=Select_focusman(username)
	if err2!=nil {
		c.JSON(500,gin.H{"status":http.StatusInternalServerError,"message":"数据库查找关注者时失败"})
	}
	c.JSON(200,gin.H{"Data":focusman_name})
	//获取具体的文章
	err3,article:=Select_article(username)
	if err3!=nil {
		c.JSON(403,gin.H{"status":http.StatusForbidden,"message":"数据库查找失败"})
	}else {
		c.JSON(200,gin.H{"status":http.StatusOK,"Data":article})
	}
	//获取具体的提问
	err4,ask:=Select_ask(username)
	if err4!=nil {
		c.JSON(403,gin.H{"status":http.StatusForbidden,"message":"数据库查询失败"})
	}else {
		c.JSON(200,gin.H{"status":http.StatusOK,"Data":ask})
	}
}

func Select_userinfor(username string)(error,Userinfor){
	var user Userinfor
	stmt,err:=db.Query("select username,question_number,focus_number from user where username=?",username)
	if err!=nil {
		fmt.Println("数据库查询失败")
	}
	defer stmt.Close()
	stmt.Scan(&user.username,&user.question_number,&user.focus_number)
	return err,user
}

func Select_focusman(username string)(error,[]name){
	var names []name
	stmt,err:=db.Query("select username,id from user where id in (select focus_id from focus where username=?)",username)
	if err!=nil {
		fmt.Println("数据库查询id时出现错误")
	}
	defer stmt.Close()
	for stmt.Next(){
		var name name
		stmt.Scan(&name.username,&name.id)
		names=append(names,name)
	}
	return err,names
}

func Select_ask(username string)(error,[]ask){
	var asks []ask
	stmt,err:=db.Query("select ask from question where username=?",username)
	if err!=nil {
		fmt.Println("查询提问时出错")
	}
	defer stmt.Close()
	for stmt.Next(){
		var ask ask
		stmt.Scan(&ask.username,&ask.ask)
		asks=append(asks,ask)
	}
	return err,asks
}

func Select_article(username string)(error,[]article){
	var articles []article
	stmt,err:=db.Query("select article from article where username=?",username)
	if err!=nil {
		fmt.Println("查询文章时出现错误")
	}
	defer stmt.Close()
	for stmt.Next(){
		var article article
		stmt.Scan(&article.username,&article.atricle)
		articles=append(articles,article)
	}
	return err,articles
}