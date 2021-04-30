package data

import "gorm.io/gorm"

var (
	DBConn *gorm.DB
)

type Category struct {
	gorm.Model
	Name     string    `json:"name"`
	Products []Product `json:"-"`
	ParentID *uint     `json:"parent_id"`
	Parent   *Category `gorm:"foreignKey:ParentID;references:ID" json:"-"`
}

type Product struct {
	gorm.Model
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"imageURL"`
	Variant     []Variant
	CategoryID  uint
	ParentID    *uint
	Parent      *Product `gorm:"foreignKey:ParentID;references:ID"`
}

type Variant struct {
	gorm.Model
	ID            uint    `json:"id" `
	Name          string  `json:"name"`
	MRP           float32 `json:"mrp"`
	DiscountPrice float32 `json:"discountPrice"`
	Size          int     `json:"size"`
	Colour        string  `json:"colour"`
	ProductID     uint
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Category{}, &Product{}, &Variant{})
	return db
}
