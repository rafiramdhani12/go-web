package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDb(){
	db, err := sql.Open("mysql", "root:@/go_products?parseTime=true")
	if err != nil{
		panic(err)
	}
	log.Println("database connected")
	DB = db
}
