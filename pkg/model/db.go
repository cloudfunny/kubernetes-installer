package model

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	dsn := "host=localhost port=5432 user=postgres password=passw0rd dbname=kubernetes-installer sslmode=disable TimeZone=Asia/Shanghai"

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to create db connection")
		panic(err)
	}

	fmt.Println("DB connect success, start to create tables...")
	CreateTables()
}

func CreateTables() {
	db.AutoMigrate(&Cluster{})
	db.AutoMigrate(&Node{})
}
