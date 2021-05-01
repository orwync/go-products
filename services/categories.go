package services

import "github.com/orwync/go-products/data"

func GetAllCategories() ([]data.Category, error) {
	var categories []data.Category
	err := data.DBConn.Preload("Products").Preload("Parent").Find(&categories).Error
	return categories, err
}

func GetCategory(id int) (data.Category, error) {
	var category data.Category
	err := data.DBConn.Preload("Products").Preload("Parent").First(&category, id).Error
	return category, err
}

func CreateCategory(category *data.Category) error {
	return data.DBConn.Create(category).Error
}

func UpdateCategory(id int, updateCategory *data.Category) error {
	var category data.Category
	err := data.DBConn.Find(&category, id).Error

	if err != nil {
		return err
	}

	err = data.DBConn.Model(&category).Updates(updateCategory).Error
	return err
}

func DeleteCategory(id int) error {
	return data.DBConn.Delete(&data.Category{}, id).Error
}
