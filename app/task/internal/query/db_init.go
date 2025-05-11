package query

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/db_todolist?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}
