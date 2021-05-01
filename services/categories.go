package services

import "github.com/orwync/go-products/data"

func GetAllCategories() ([]data.Category, error) {
	var categories []data.Category
	err := data.DBConn.Preload("Products").Find(&categories).Error
	return categories, err
}

func GetCategory(id int) (data.Category, error) {
	var category data.Category
	err := data.DBConn.Preload("Products").First(&category, id).Error
	return category, err
}
