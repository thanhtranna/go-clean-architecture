package datastore

import (
	"log"

	"golang-clean-architecture/config"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewDB() *gorm.DB {
	DBMS := "mysql"
	mySqlConfig := &mysql.Config{
		User:                 config.Config.Database.User,
		Passwd:               config.Config.Database.Password,
		Net:                  config.Config.Database.Net,
		Addr:                 config.Config.Database.Addr,
		DBName:               config.Config.Database.DBName,
		AllowNativePasswords: config.Config.Database.AllowNativePasswords,
		Params: map[string]string{
			"parseTime": config.Config.Database.Params.ParseTime,
		},
	}

	db, err := gorm.Open(DBMS, mySqlConfig.FormatDSN())

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
