package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

var db *gorm.DB
/*const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "root"
	dbname   = "propertiesWeb"
)*/

func init()  {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	user := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbname := os.Getenv("db_name")
	host := os.Getenv("db_host")
	port := os.Getenv("db_port")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	conn, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	db = conn

	db.Debug().AutoMigrate(&Account{},&Property{},&Rating{},&Visit{})
}

func GetDB() *gorm.DB {
	return db
}