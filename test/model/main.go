package model

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB


func init() {
	db, _ = sql.Open("mysql", "root:lhr20001221@tcp(localhost:3306)/redrock?charset=utf8")
}
func DBConn() *sql.DB{
	return  db
}

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
