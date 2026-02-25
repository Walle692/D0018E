package products

type SellerProduct struct {
	Product_id   int     `json:"product_id"`
	Product_name string  `json:"product_name"`
	Manufacturer string  `json:"manufacturer"`
	Picture_url  string  `json:"picture_url"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	Active       bool    `json:"active"`
	Description  string  `json:"description"`
	Screen_size  float32 `json:"screen_size"`
}
