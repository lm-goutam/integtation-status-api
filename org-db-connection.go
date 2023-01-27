package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var urIDSN = "root:admin@123@tcp(localhost:3030)/intg_stat"
var err error

func DataMigration() {

	Database, err = gorm.Open(mysql.Open(urIDSN), &gorm.Config{})
	if err != nil {
		fmt.Print(err.Error())
		panic("Connection Failed :(")
	}
	Database.AutoMigrate(&Org{})
}
