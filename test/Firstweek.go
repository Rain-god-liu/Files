package main

import(
	"github.com/jinzhu/gorm"
	"log"
)

var(
	DB *gorm.DB
)


func init(){
	mysql, err := gorm.Open("mysql", "root:lhr20001221@(127.0.0.1:3306)/firstweek?charset=utf8&parseTime=true")
    if err!=nil{
      log.Fatal(err.Error())
	}
	DB=mysql
	defer mysql.Close()
}
