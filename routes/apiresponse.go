package routes

import "time"

// a category returns in the response
// swagger:response categoryResponse
type Category struct {
	// Single category
	// in:body
	ID        int
	Name      string     `json:"name"`
	Products  []Product  `json:"products"`
	ParentID  *uint      `json:"parentId"`
	Parent    []Category `json:"parent"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// A list of categoriest returns in the response
// swagger:response categoriesResponse
type CategoriesResponse struct {
	// All the categoriest in the system
	// in:body
	Body []Category
}

// swagger:parameters createCategory editCategory
type CreateCategoryResponse struct {
	// category name
	// in: query
	// required: true
	Name string `json:"name"`
	// category parent id
	ParentID *uint `json:"parentId"`
}

// swagger:parameters getCategory editCategory deleteCategory
type CategoryIDparameterWrapper struct {
	// The id of the category
	// in:path
	// required: true
	ID int `json:"id"`
}

// a product returns in the response
// swagger:response productResponse
type Product struct {
	// Single product
	// in:body
	ID          int
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageURL    string    `json:"imageURL"`
	Variant     []Variant `json:"variants"`
	CategoryID  uint      `json:"categoryId"`
	ParentID    uint      `json:"parentId"`
	Parent      []Product `json:"parent"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

// A list of product returns in the response
// swagger:response productsResponse
type ProductsResponse struct {
	// All the product in the system
	// in:body
	Body []Product
}

// swagger:parameters createProduct editProduct
type CreateProductResponse struct {
	// The name of the product
	// in:query
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"imageURL"`
	CategoryID  uint   `json:"categoryId"`
	ParentID    uint   `json:"parentId"`
}

// swagger:parameters getProduct editProduct deleteProduct
type ProductIDparameterWrapper struct {
	// The id of the product
	// in:path
	// required: true
	ID int `json:"id"`
}

// a variant returns in the response
// swagger:response variantResponse
type Variant struct {
	// Single variant
	// in:body
	Name          string  `json:"name"`
	MRP           float32 `json:"mrp"`
	DiscountPrice float32 `json:"discountPrice"`
	Size          int     `json:"size"`
	Colour        string  `json:"colour"`
	ProductID     uint    `json:"productId"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

// A list of variants returns in the response
// swagger:response variantsResponse
type VariantsResponse struct {
	// All the variants in the system
	// in:body
	Body []Variant
}

// swagger:parameters createVariant editVairant
type CreateVariantResponse struct {
	// The details of the varinat
	// in:query
	Name          string  `json:"name"`
	MRP           float32 `json:"mrp"`
	DiscountPrice float32 `json:"discountPrice"`
	Size          int     `json:"size"`
	Colour        string  `json:"colour"`
	ProductID     uint    `json:"productId"`
}

// swagger:parameters getVariant editVariant deleteVariant
type VariantIDparameterWrapper struct {
	// The id of the varinat
	// in:path
	// required: true
	ID int `json:"id"`
}

// Status response
// swagger:model statusMessage
type StatusMessage struct {
	// Status messages
	// in:body
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
