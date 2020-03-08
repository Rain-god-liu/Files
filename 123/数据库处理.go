package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)



func main() {
	db, err := sql.Open("mysql", "root:lhr20001221@tcp(localhost:3306)/user?charset=utf8")


	if err != nil {
		log.Fatal(err)
	}


	//增删改查请在下面选。
		//insertDB(db)
		//deleteDB(db)
		//updateDB(db)
		//selectDB(db)
}

	func insertDB(db * sql.DB) {
		stmt, err := db.Prepare("insert into test(name, age) values (?, ?)")
		defer stmt.Close()
		if err != nil {
			log.Fatal(err)
		}

		stmt.Exec("奇点", "18")

	}

	func deleteDB(db * sql.DB) {
		stmt, err := db.Prepare("delete from test where name='奇点'")
		defer stmt.Close()
		if err != nil {
			log.Fatal(err);
		}

		stmt.Exec();
	}

	func updateDB(db * sql.DB) {
		stmt, err := db.Prepare("UPDATE test SET name = '隔壁老王' WHERE name='奇点'")
		if err != nil {
			log.Fatal(err)
		}
		stmt.Exec();
	}

	func selectDB(db * sql.DB) {
		stmt, err := db.Query("select * from test;")
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
