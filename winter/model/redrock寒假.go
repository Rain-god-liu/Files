package model

import (
	"database/sql"
)

var db *sql.DB

func DBConn() *sql.DB{
	return  db
}

