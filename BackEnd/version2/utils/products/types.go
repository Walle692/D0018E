package products

type Product struct {
	Product_id     int     `json:"product_id"`
	Product_name   string  `json:"product_name"`
	Manufacturer   string  `json:"manufacturer"`
	Seller_user_id int     `json:"seller_user_id"`
	Picture_url    string  `json:"picture_url"`
	Price          float64 `json:"price"`
	Stock          int     `json:"stock"`
	Active         bool    `json:"active"`
	Description    string  `json:"description"`
	Screen_size    float64 `json:"screen_size"`
}
