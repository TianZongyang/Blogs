package utils

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

func InitDB(dialect string, connStr string, maxPoolSize, minPoolSize int) *gorm.DB {
	db, err := gorm.Open(dialect, connStr)
	if err != nil {
		log.Fatalln(err)
	}
	db.SingularTable(true)
	db.DB().SetMaxOpenConns(maxPoolSize)
	db.DB().SetMaxIdleConns(minPoolSize)
	db.DB().SetConnMaxLifetime(10 * time.Minute)
	return db
}

func CloseDB(gdb *gorm.DB) {
	err := gdb.Close()
	if err != nil {
		log.Println("close db connection error !")
	} else {
		log.Println("success close db connection !")
	}
}
