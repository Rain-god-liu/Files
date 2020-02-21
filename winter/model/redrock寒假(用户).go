package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)



var account=make(map[string]string)

func UserSignup(username string,password string)bool{
	stmt,err:=DBConn().Prepare(
		"insert into user(username,password)values(?,?) ")
	if err!=nil{
		fmt.Println("fail to insert")
		return false
	}
	defer stmt.Close()
	_,err=stmt.Exec(username,password)
	if err!=nil{
		fmt.Println("fail to insert")
		return false
	}
	return false
}

func UserSignin( username string,password string)bool{
	stmt,err:=db.Query("select password from user where username=?",username)
	if err!=nil{
		log.Fatal(err)
		return false
	}
	defer stmt.Close()
	for stmt.Next() {
		var row string
		err = stmt.Scan(&row)
		if row==password{
			return true
		}
	}
	return false
}



func Checkuser(username string)bool{
	stmt,err:=db.Query("select username from user ")
	defer stmt.Close()
	var name string
	if err!=nil {
		fmt.Println(err)
		return false
	}
	for stmt.Next(){
		err:=stmt.Scan(&name)
		if err!=nil{
			fmt.Println(err)
			return false
		}
		account[name]=name
	}
	if username!=account[name] {
		return true
	}else {
		return false
	}
}


func Cancellation(c *gin.Context){
	username,err:=c.Cookie("username")
	if err!=nil {
		fmt.Println("没有找到cookie")
	}
	c.SetCookie("username",username,-1, "/", "localhost", false, true)
}

func Focus_other(c *gin.Context){
	//直接从请求连接里找到
	focus_name:=c.PostForm("username")
	username,err:=c.Cookie("username")
	if err!=nil {
		fmt.Println("cookie查询出现错误")
	}
	//查询被关注者的id
	focus_id:=Select_focusid(focus_name)
	if Focus_sql(username,focus_id) {
		c.JSON(200,gin.H{"status":http.StatusOK,"message":"关注成功"})
	}else {
		c.JSON(500,gin.H{"status":http.StatusInternalServerError,"message":"数据库Insert报错"})
	}
}

func Select_focusid(focus_name string)(int){
	var focus_id int
	stmt,err:=db.Query("select id from user where username=?",focus_name)
	if err!=nil {
		fmt.Println("查询focus_id时出现错误")
	}
	defer stmt.Close()
	stmt.Scan(&focus_id)
	return focus_id
}

func Focus_sql(username string,focus_id int)bool{
	stmt,err:=DBConn().Prepare("insert into focus(username,focus_id) values(?,?)")
	if err!=nil {
		fmt.Println("数据库插入关注者是出现错误")
		return false
	}
	defer stmt.Close()
	_,err=stmt.Exec(username,focus_id)
	return true
}