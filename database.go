package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	// for mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	// DB - shared conn
	DB *gorm.DB
)

func init() {
	var err error
	if DB, err = InitDB(); err != nil {
		panic(fmt.Sprintf("No error should happen when connecting to test database, but got err=%+v", err))
	}

}

// InitDB -
func InitDB() (db *gorm.DB, err error) {
	addr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=true",
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
	)

	// addr := fmt.Sprintf("postgresql://%v:%v@localhost:26257/%v?sslmode=disable",
	// 	os.Getenv("DATABASE_USERNAME"),
	// 	os.Getenv("DATABASE_PASSWORD"),
	// 	os.Getenv("DATABASE_NAME"),
	// )

	db, errConn := gorm.Open("mysql", addr)
	// db, errConn := gorm.Open("postgres", addr)
	if errConn != nil {
		log.Fatal(errConn)
	}
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)

	return
}
