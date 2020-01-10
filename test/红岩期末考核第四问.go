package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
)

var db *sql.DB
var username string
var administratorname string
var id int
var administratorid int
var jifen int

func init() {
	db, _ = sql.Open("mysql", "root:lhr20001221@tcp(localhost:3306)/redrock?charset=utf8")

}
func DBConn() *sql.DB{
	return  db
}

func main() {
	r := gin.Default()
	r.POST("/register", Registe)
	r.POST("/login", Login)
	r.POST("/insertprize", Insert)
	r.POST("/update", Update)
	r.POST("/select", Select)
	r.Run(":8080")
}

func Insert(c *gin.Context){
	prizename:=c.PostForm("prizename")
	prizecost:=c.PostForm("prizecost")
	fmt.Println("prize"+prizename+prizecost)
	if insert(prizename,prizecost){
		c.JSON(500,gin.H{"status":http.StatusInternalServerError,"message":"数据库Insert报错"})
	}else {
		c.JSON(200, gin.H{"status": http.StatusOK, "message": "添加成功"})
	}
}

func Update(c *gin.Context) {
	prizename:=c.PostForm("prizename")
	prizecost:=c.PostForm("prizecost")
	update(prizename,prizecost)


}

func Select(c *gin.Context){

}

func Registe(c *gin.Context){
	username:=c.PostForm("username")
	password:=c.PostForm("password")
	fmt.Println("user:"+username+password)
	if UserSignup(username,password){
		c.JSON(500,gin.H{"status":http.StatusInternalServerError,"message":"数据库Insert报错"})
	}else {
		c.JSON(200, gin.H{"status": http.StatusOK, "message": "注册成功"})
	}
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if UserSignin(username, password) {
		c.SetCookie("username", username, 10, "localhost:8080", "localhost", false, true)
		c.JSON(200, gin.H{"status": http.StatusOK, "message": "登录成功"})
	} else {
		c.JSON(403, gin.H{"status": http.StatusForbidden, "message": "登录失败，用户名或密码错误"})
	}
	stmt,err:= db.Prepare("update jifen ='?+1' where username ='?'")
	defer stmt.Close()
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

}


func UserSignup(username string,password string)bool{
	stmt,err:=DBConn().Prepare(
		"insert into userlist(username,password)values(?,?) ")
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
	stmt,err:=db.Query("select password from userlist where username=?",username)
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

//判断该用户是不是管理员
func Guanliyuan(db * sql.DB) {
	stmt, err := db.Query("select id from userlist where username ='lhr'")
	defer stmt.Close()
	if err != nil {
		log.Fatal(err)
	}
	for stmt.Next() {
		err := stmt.Scan(&administratorid)
		if err != nil {
			log.Fatal(err)
		}

	}
}


//管理员权限目录，对奖品进行添加
func insert(prizename string,prizecost string)bool {
	Guanliyuan(db)
	if id==administratorid {
		fmt.Println("欢迎您，管理员")
	}else {
		log.Fatal("对不起，您不是管理员")
	}
	stmt, err := DBConn().Prepare("insert into prizelist (prizename, prizecost) values (?, ?)")
	defer stmt.Close()
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_,err=stmt.Exec(prizename,prizecost)
	if err!=nil{
		fmt.Println("fail to insert")
		return false
	}

	return false


}

//对奖品进行删除
func delete(db * sql.DB) {
	Guanliyuan(db)
	if id==administratorid {
		fmt.Println("欢迎您，管理员")
	}else {
		log.Fatal("对不起，您不是管理员")
	}
	stmt, err := DBConn().Prepare("delete from prizelist where prizename='?'")
	defer stmt.Close()
	if err != nil {
		log.Fatal(err);
	}

	stmt.Exec();
}

//对奖品进行更新
func update(prizename string,prizecost string) {
	Guanliyuan(db)
	if id==administratorid {
		fmt.Println("欢迎您，管理员")
	}else {
		log.Fatal("对不起，您不是管理员")
	}
	stmt, err := DBConn().Prepare("UPDATE prizelist SET prizecost = '?' WHERE prizename='?'")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec();
}

//查看用户积分
func select_(db * sql.DB) {
	Guanliyuan(db)
	if id==administratorid {
		fmt.Println("欢迎您，管理员")
	}else {
		log.Fatal("对不起，您不是管理员")
	}
	stmt, err := db.Query("select * from signlist;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for stmt.Next() {
		var age int
		var name string

		err := stmt.Scan(&name, &age)
		if err != nil {

			log.Fatal(err)
		}
		fmt.Println(name, age)
	}

}
