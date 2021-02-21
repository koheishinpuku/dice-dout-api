package db

import (
	"database/sql"
	"fmt"
	"os"
)

var host = os.Getenv("DB_HOST")
var port = os.Getenv("DB_PORT")
var user = os.Getenv("DB_USER")
var pwd = os.Getenv("DB_PWD")
var name = os.Getenv("DB_NAME")

// DB : db connection
var DB *gorm.DB

var mySQLDB *sql.DB

func Init() {
	connString := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pwd, host, port, name)
	DB, err = gorm.Open("mysql", connString)
}
