package model

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

var uid int

func init(){
	db, _ = sql.Open("mysql", "root:lhr20001221@tcp(localhost:3306)/redrock?charset=utf8")
}

type Liuyan struct {
	Child_message *[]Liuyan
	Username string
	Message string
	id int
	uid int
}

func SendMessage(username string,message string)bool{
	stmt,err:=DBConn().Prepare(
		"insert into message(username,message)values(?,?)")
	if err!=nil{
		fmt.Println("fail to insert the message")
		return false
	}
	defer stmt.Close()
	_,err=stmt.Exec(username,message)
	if err!=nil{
		fmt.Println("fail to insert")
		return false
	}
	return true
}

func GetMesage(id int,maxtime int)(error, []Liuyan){
	var message []Liuyan
	stmt,err:=db.Query("select username,message,uid,id from message where id=?",id)
	if err!=nil{
		fmt.Println("没有这条留言")
	}
	defer stmt.Close()
	for stmt.Next() {
		var msg Liuyan
		stmt.Scan(&msg.Username, &msg.Message,&msg.uid,&msg.id)
		if err!=nil{
			fmt.Println("fail to do this")
		}
		err,child:=GetMesage(msg.uid,10)
		if err!=nil{
			fmt.Println("发生错误")
		}
		msg.Child_message=&child
		message=append(message,msg)
	}
	return err,message
}



//追加留言
func Find_main_id(c *gin.Context){
	var messages Liuyan
	username:=c.PostForm("username")
	message:=c.PostForm("message")
	stmt,err:=db.Query("select id,uid from message where message=? and username=?",message,username)
	if err!=nil {
		log.Fatal("没有找到你要追加的主留言")
	}
	stmt.Scan(&messages.id,&messages.uid)
	uid=messages.uid
	uidNew:=messages.id
	Update_main_uid(uidNew,uid)
}

func Update_main_uid(uid int,id int){
	stmt,err:=db.Prepare("update message set uid='?'value(?) where id='?'value(?)")
	if err!=nil {
		fmt.Println("更改uid时出现数据库错误")
	}
	defer stmt.Close()
	stmt.Exec(uid,id)
}

func Append(username string,message string,uid int){
	stmt,err:=DBConn().Prepare("insert message(username,message,uid)values(?,?,?) ")
	if err!=nil {
		log.Fatal("留言发表失败")
	}
	defer stmt.Close()
	stmt.Exec(username,message,uid)
}

