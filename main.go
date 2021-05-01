package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/orwync/go-products/data"
	"github.com/orwync/go-products/helper"
	"github.com/orwync/go-products/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDatabase() {
	connectionString := "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(connectionString, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_URL"), os.Getenv("DB_NAME"))
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
	godotenv.Load()
	initDatabase()
	r := routes.HandleRoutes()

	r.Use(helper.SetContentType)

	s := &http.Server{
		Addr:         ":8000",
		Handler:      r,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	s.ListenAndServe()
}
