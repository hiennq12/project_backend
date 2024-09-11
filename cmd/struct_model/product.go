package struct_model

import "time"

type Product struct {
	Id            int       `json:"id"`
	UserId        int       `json:"user_id"`
	CategoryId    int       `json:"category_id"`
	ProductName   string    `json:"product_name"`
	Description   string    `json:"description"`
	Price         float64   `json:"price"`
	Condition     string    `json:"condition"` // New, Used, Refurbished
	Location      string    `json:"location"`
	StockQuantity int       `json:"stock_quantity"`
	Weight        float64   `json:"weight"`
	Dimensions    string    `json:"dimensions"`
	SKU           string    `json:"sku"`
	Brand         string    `json:"brand"`
	Warranty      string    `json:"warranty"`
	IsNegotiable  bool      `json:"is_negotiable"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	IsActive      bool      `json:"is_active"`
}

type ProductsRequest struct {
	ProductId  int64   `json:"product_id"`
	ProductIds []int64 `json:"product_ids"`
	Limit      uint64  `json:"limit"`
}

type InsertProductsRequest struct {
	UserId        int       `json:"user_id,omitempty"`
	CategoryId    int       `json:"category_id,omitempty"`
	ProductName   string    `json:"product_name,omitempty"`
	Description   string    `json:"description,omitempty"`
	Price         float64   `json:"price,omitempty"`
	Condition     string    `json:"condition,omitempty"`
	Location      string    `json:"location,omitempty"`
	StockQuantity int       `json:"stock_quantity,omitempty"`
	Weight        float64   `json:"weight,omitempty"`
	Dimensions    string    `json:"dimensions,omitempty"`
	SKU           string    `json:"sku,omitempty"`
	Brand         string    `json:"brand,omitempty"`
	Warranty      string    `json:"warranty,omitempty"`
	IsNegotiable  bool      `json:"is_negotiable,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
	IsActive      bool      `json:"is_active,omitempty"`
}
