package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/michaelwp/golang-gorm/helpers"
	"github.com/michaelwp/golang-gorm/models"
	"log"
)

func MySql() *gorm.DB{
	DbUser := helpers.GetEnv("DB_USER")
	DbPass := helpers.GetEnv("DB_PASS")
	DbHost := helpers.GetEnv("DB_HOST")
	DbPort := helpers.GetEnv("DB_PORT")
	DbName := helpers.GetEnv("DB_NAME")
	DbUri := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DbUser, DbPass, DbHost, DbPort, DbName)

	dbMysql, err := gorm.Open("mysql", DbUri)
	if err != nil {log.Fatal(err)}

	dbMysql.AutoMigrate(&models.User{}, &models.Credential{})

	fmt.Println("connected to mysql database")

	return dbMysql
}
