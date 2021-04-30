package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/orwync/go-products/data"
	"github.com/orwync/go-products/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDatabase() {
	dsn := "root:root@tcp(127.0.0.1:3306)/productsdb?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	data.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to DB")
	}

	fmt.Println("Database connection sucessfully opened")

	data.DBMigrate(data.DBConn)
	fmt.Println("Database migrated")
}

func main() {
	initDatabase()
	r := routes.HandleRoutes()

	s := &http.Server{
		Addr:         ":8000",
		Handler:      r,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	s.ListenAndServe()
}
