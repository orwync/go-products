package services

import (
	"errors"

	"github.com/orwync/go-products/data"
)

func GetAllProducts() ([]data.Product, error) {
	var products []data.Product
	err := data.DBConn.Preload("Variant").Preload("Parent").Find(&products).Error
	return products, err
}

func GetProduct(id int) (data.Product, error) {
	var product data.Product
	err := data.DBConn.Preload("Variant").Preload("Parent").First(&product, id).Error
	return product, err
}

func CreateProduct(product *data.Product) error {
	return data.DBConn.Create(product).Error
}

func UpdateProduct(id int, updateProduct *data.Product) error {
	var product data.Product
	data.DBConn.Find(&product, id)

	if product.ID == 0 {
		return errors.New("record not found")
	}

	return data.DBConn.Model(&product).Updates(updateProduct).Error
}

func DeleteProduct(id int) error {
	return data.DBConn.Delete(&data.Product{}, id).Error
}
