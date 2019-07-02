package config 

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "github.com/go-sql-driver/mysql"
)

func GetDB() (*gorm.DB, error) {
	dbDriver := "mysql"
	dbName := "demo"
	dbUser := "root"
	dbPassword := ""

	db, err := gorm.Open(dbDriver, dbUser+":"+dbPassword+"@/"+dbName+"?charset=utf8&parseTime=True")
	if err !=nil{
		return nil,err
	}
	return db, nil

}
