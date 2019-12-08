package model

import (
	"fmt"

)

type Liuyan struct {
	Username string
	Message string
}



func SendMessage(username string,message string)bool{
	stmt,err:=DBConn().Prepare(
		"insert into user2(username,message)values(?,?)")
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



func GetMesage(maxtime int)(error, []Liuyan){
	var message []Liuyan
       stmt,err:=db.Query("select username,message from user2 where username=?","lhr")
        if err!=nil{
	fmt.Println("没有这条留言")
           }
	defer stmt.Close()
	var msg Liuyan
       for stmt.Next() {
		   stmt.Scan(&msg.Username, &msg.Message)
	   }
	   message=append(message,msg)
	   if err!=nil{
	   	fmt.Println("fail to do this")
	   }
	return err,message
}
