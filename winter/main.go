package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
	"strings"
	"winter/model"
)

var db *sql.DB
var check=make(map[string]string)

func DBConn() *sql.DB{
	return  db
}

func main(){
	db, _ = sql.Open("mysql", "root:lhr20001221@tcp(localhost:3306)/redrock?charset=utf8")
	r:=gin.Default()
	r.Use(Cors())
	r.POST("/register",Panduandenglu)
	r.POST("/login",Panduandenglu)
	r.POST("/cancellation",model.Cancellation)
	r.POST("/sendmasge",SendMsg)
	r.GET("/getmasge",GetMsg)
	r.POST("/insert",Send_information)
	r.POST("/search",model.Search)
	r.GET("/userinfor",model.User_information)
	r.POST("/user",model.User_information)
	r.POST("/ask",model.Ask)
	r.POST("/focus",model.Focus_other)
	r.Run(":8080")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method      //请求方法
		origin := c.Request.Header.Get("Origin")        //请求头部
		var headerKeys []string                             // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")        // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")      //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")      // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")        // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")       //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")       // 设置返回格式是json
		}
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next()        //  处理请求
	}
}

func Panduandenglu(c *gin.Context){
	username:=c.PostForm("username")
	password:=c.PostForm("password")
	if Makesure_name(username,password) {
		Login(c,username,password)
	}else {
		Registe(c,username,password)
	}
}

func Makesure_name(username string,password string)bool{
	stmt,err:=db.Query("select username from user ")
	defer stmt.Close()
	if err!=nil {
		log.Fatal("查询出现错误")
		return false
	}
	var name string
	for stmt.Next(){
		err:=stmt.Scan(&name)
		if err!=nil{
			log.Fatal(err)
			return false
		}
		check[name]=name
	}
	if username!=check[name]{
		return false
	}
	return true
}

func Registe(c *gin.Context,username string,password string){
	fmt.Println("user:"+username+password)
	if model.Checkuser(username){
		if model.UserSignup(username,password){
			c.JSON(500,gin.H{"status":http.StatusInternalServerError,"message":"数据库Insert报错"})
		}else {
			c.JSON(200, gin.H{"status": http.StatusOK, "message": "注册成功"})
		}}else {
		c.JSON(200,gin.H{"status":http.StatusOK,"message":"注册失败，该用户名已被注册"})
	}
}

func Login(c *gin.Context,username string,password string) {
	if model.UserSignin(username, password) {
		c.SetCookie("username", username, 3600, "/", "localhost", false, true)
		c.JSON(200, gin.H{"status": http.StatusOK, "message": "登录成功"})
	} else {
		c.JSON(403, gin.H{"status": http.StatusForbidden, "message": "登录失败，用户名或密码错误"})
	}
}

func Send_information(c *gin.Context){
	keyword:=c.PostForm("keyword")
	information:=c.PostForm("information")
	if model.Insert_information(keyword,information) {
		c.JSON(500,gin.H{"status":http.StatusInternalServerError,"message":"数据库Insert报错"})
	}
}

func SendMsg(c *gin.Context){
	username,_:=c.Cookie("username")
	fmt.Println("username"+username)
	/*if err != nil{
		c.JSON(500,gin.H{"status": http.StatusForbidden,"message":"cookie读取失败"})
		return
	}*/
	message:=c.PostForm("message")
	if model.SendMessage(username,message){
		c.JSON(200, gin.H{"用户名":username,"留言内容":message})
	}else {
		c.JSON(403, gin.H{"status": http.StatusForbidden, "message": "发送失败"})
	}
}

func GetMsg(c *gin.Context){
	id:=c.PostForm("id")
	idNew, _ :=strconv.Atoi(id)
	err,messages:=model.GetMesage(idNew,10)
	if err !=nil{
		c.JSON(500,gin.H{"status": http.StatusInternalServerError,"message":"数据库读取失败"})
	}
	c.JSON(200,gin.H{"Data":messages})
}
