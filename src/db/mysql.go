package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type MySQL struct {
	Url string
	DB  *gorm.DB
}

func (mysql *MySQL) InitMySQL() {
	db, err := gorm.Open("mysql", mysql.Url)
	if err != nil {
		log.Printf("open mysql error(%v)", err)
		panic(err)
	}
	mysql.DB = db
}

func (mysql *MySQL) CloseMySQL() {
	err := mysql.DB.Close()
	if err != nil {
		log.Printf("close mysql error(%v)", err)
		panic(err)
	}
}
