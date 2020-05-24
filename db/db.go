package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"golang-gorm/models"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func MySql() *gorm.DB{
	DbUser := loadDBCred("DB_USER")
	DbPass := loadDBCred("DB_PASS")
	DbHost := loadDBCred("DB_HOST")
	DbPort := loadDBCred("DB_PORT")
	DbName := loadDBCred("DB_NAME")
	DbUri := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DbUser, DbPass, DbHost, DbPort, DbName)

	dbMysql, err := gorm.Open("mysql", DbUri)
	if err != nil {log.Fatal(err)}

	dbMysql.AutoMigrate(&models.User{}, &models.Credential{})

	fmt.Println("connected to mysql database")

	return dbMysql
}

func loadDBCred(key string) string {
	err := godotenv.Load(".env")
	if err != nil {log.Fatal(err)}

	return os.Getenv(key)
}
