package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/jinzhu/gorm"
	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var host = os.Getenv("DB_HOST")
var port = os.Getenv("DB_PORT")
var user = os.Getenv("DB_USER")
var pwd = os.Getenv("DB_PWD")
var name = os.Getenv("DB_NAME")

// DB : db connection
var DB *gorm.DB

var mySQLDB *sql.DB

// Init :
func Init() *sql.DB {
	connString := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pwd, host, port, name)
	DB, err := gorm.Open("mysql", connString)
	if err != nil {
		fmt.Println("DB_CONNECT_ERROR")
		panic(err)
	}
	mySQLDB, _ = sql.Open("mysql", connString)

	DB.LogMode(true)
	fmt.Println("DB_CONNECT!!")

	return mySQLDB
}

func main() {
	args := os.Args

	for _, command := range args[1:] {
		switch command {
		case "migrate":
			migrateDatabase()
		case "rollback":
			rollbackDatabase()
		}
	}
}

func migrateDatabase() {
	fmt.Println("[Migrate DB]")

	mySQLDB = Init()

	driver, err := mysql.WithInstance(mySQLDB, &mysql.Config{})
	if err != nil {
		log.Fatal("mysql.WithInstance:", err)
		return
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"mysql",
		driver,
	)

	if err != nil {
		log.Fatal("migrate.NewWithDatabaseInstance:", err)
	}

	if err = m.Up(); err != nil {
		log.Fatal("m.Up:", err)
	}
}

func rollbackDatabase() {
	fmt.Println("[Rollback DB]")

	mySQLDB = Init()
	driver, err := mysql.WithInstance(mySQLDB, &mysql.Config{})
	if err != nil {
		log.Fatal("mysql.WithInstance:", err)
		return
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"mysql",
		driver,
	)

	if err != nil {
		log.Fatal("migrate.NewWithDatabaseInstance:", err)
	}

	if err = m.Down(); err != nil {
		log.Fatal("m.Up:", err)
	}
}
