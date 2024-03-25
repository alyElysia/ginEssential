package common

import (
	"fmt"
	"ginEssential/model"
	"github.com/jinzhu/gorm"
)

//var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "ginEssential"
	username := "root"
	password := "alyElysia"
	charset := "utf8"
	loc := "Local"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username, password, host, port, database, charset, loc,
	)

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("数据库连接失败:" + err.Error())
	}
	db.AutoMigrate(&model.User{})
	return db
}

//func GetDB() *gorm.DB {
//
//}
