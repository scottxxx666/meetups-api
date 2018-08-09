package app

import (
	"os"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var err error

func init() {
	DB, err = gorm.Open("mysql", os.Getenv("DB_ACCOUNT")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_DATABASE")+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}
