package data

import (
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

type Category struct {
	gorm.Model
	Name     string    `json:"name"`
	Products []Product `json:"products"`
	ParentID *uint     `json:"parentId"`
	Parent   *Category `gorm:"foreignKey:ParentID;references:ID" json:"-"`
}

type Product struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageURL    string    `json:"imageURL"`
	Variant     []Variant `json:"variants"`
	CategoryID  uint      `json:"categoryId"`
	ParentID    *uint     `json:"parentId"`
	Parent      *Product  `gorm:"foreignKey:ParentID;references:ID" json:"-"`
}

type Variant struct {
	gorm.Model
	Name          string  `json:"name"`
	MRP           float32 `json:"mrp"`
	DiscountPrice float32 `json:"discountPrice"`
	Size          int     `json:"size"`
	Colour        string  `json:"colour"`
	ProductID     uint    `json:"productId"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Category{}, &Product{}, &Variant{})
	return db
}
