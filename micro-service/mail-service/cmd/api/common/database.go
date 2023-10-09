package common

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() *gorm.DB {
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port, _ := strconv.Atoi(viper.GetString("database.port"))
	dbName := viper.GetString("database.dbName")
	charset := viper.GetString("database.charset")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		dbName,
		charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
	return db
}

func GetDB() *gorm.DB {

	return DB
}
