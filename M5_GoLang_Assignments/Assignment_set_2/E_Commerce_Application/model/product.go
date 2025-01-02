package model

// Product represents an item in the e-commerce platform.
type Product struct {
    ID          int     `json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
    Stock       int     `json:"stock"`
    CategoryID  int     `json:"category_id"`
}
