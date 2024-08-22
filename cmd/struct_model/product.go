package struct_model

import "time"

type Product struct {
	ProductID     int       `json:"product_id"`
	UserID        int       `json:"user_id"`
	CategoryID    int       `json:"category_id"`
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
