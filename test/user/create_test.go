package user

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
)



func TestLogin(t *testing.T) {
	DB.HasTable(user{})
	DB.CreateTable(&user{})
}

