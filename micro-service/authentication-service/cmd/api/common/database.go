package common

import (
	"authentication/cmd/api/model"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var DB *gorm.DB

// connect database
func InitDB() *gorm.DB {
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		host, username, password, database, port)

	db, err := gorm.Open(driverName, dsn)
	if err != nil {
		panic("failed to connect postgre, err: " + err.Error())
	} else {
		log.Println("connect to postgre")
	}
	db = db.AutoMigrate(&model.User{})
	DB = db
	return db

}

func GetDB() *gorm.DB {
	return DB
}
