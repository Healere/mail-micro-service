package common

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestAll(t *testing.T) {
	t.Run("testing insert data", TestInsert)
}

type Mailer struct {
	username   string
	password   string
	name       string
	maildir    string
	quota      int
	local_part string
	domain     string
	created    string
	modified   string
	active     int
}

func TestInsert(t *testing.T) {
	db := GetConnect()
	mailer := []interface{}{}
	mailer = append(mailer,
		"user1@lhjpost.top",
		"user1@lhjpost.top",
		"user1",
		"lhjpost.top/user1/",
		0,
		"lhjpost",
		"lhjpost.top",
		time.Now().Format("2006-01-02 15:04:05"),
		time.Now().Format("2006-01-02 15:04:05"),
		1,
	)
	fmt.Println(mailer...)
	sql := "insert into mailbox values"
	sql += "(?,?,?,?,?,?,?,?,?,?)"
	sql = strings.TrimSuffix(sql, ",")
	fmt.Println(sql)
	tx := db.Begin()
	if err := tx.Exec(sql, mailer...).Error; err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	//db.Table("mailbox").Debug().Create(&mailer)
}

func GetConnect() *gorm.DB {
	//workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath("../config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

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
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("connect error: " + err.Error())
	}

	return db
}
