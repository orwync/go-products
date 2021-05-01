package services

import (
	"errors"

	"github.com/orwync/go-products/data"
)

func GetAllVariants() ([]data.Variant, error) {
	var variants []data.Variant
	err := data.DBConn.Find(&variants).Error
	return variants, err
}

func GetVariant(id int) (data.Variant, error) {
	var variant data.Variant
	err := data.DBConn.First(&variant, id).Error
	return variant, err
}

func CreateVariant(variant *data.Variant) error {
	return data.DBConn.Create(variant).Error
}

func UpdateVariant(id int, updateVariant *data.Variant) error {
	var variant data.Variant
	data.DBConn.Find(&variant, id)

	if variant.ID == 0 {
		return errors.New("record not found")
	}

	return data.DBConn.Model(&variant).Updates(updateVariant).Error
}

func DeleteVariant(id int) error {
	return data.DBConn.Delete(&data.Variant{}, id).Error
}
