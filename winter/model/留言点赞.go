package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

var zan int

func Dianzan(c *gin.Context){
	//获取cookie里的username
	username,err:=c.Cookie("username")
	if err!=nil {
		fmt.Println("读取cookie时出现错误")
	}
	id:=Getuserid(username)
	//输入1表示觉得很赞
	view:=c.PostForm("view")
	viewNew,_:=strconv.Atoi(view)
	Dianzan_main(id,viewNew)
}

func Dianzan_main(id int,view int) {
	//获取点赞
	zan=Select_zan(id)
	if Dianzan_user(view) {
		zanNew := zan + 1
		stmt, err := db.Prepare("update message set zan='?' where id='?'values(?,?)")

		if err != nil {
			fmt.Println("更改点赞数时出现错误")
		}

		stmt.Exec(zanNew,id)
		defer stmt.Close()
	}else {

	}

}

func Dianzan_user(view int)bool{
	if view!=1 {
		return false
	}
	return true
}

func Getuserid(username string)(int){
	//直接读取cookie里的username,然后去找对应的username的id
	var id int
	stmt,err:=db.Query("select id from user where username=?",username)
	if err!=nil {
		fmt.Println("数据库查找id时发生错误")
	}
	defer stmt.Close()
	stmt.Scan(&id)
	return id
}

func Select_zan(id int)(int){
	stmt,err:=db.Query("select zan from message where id=?",id)
	if err!=nil {
		fmt.Println("数据库查找id时发生错误")
	}
	defer stmt.Close()
	stmt.Scan(&zan)
	return zan
}