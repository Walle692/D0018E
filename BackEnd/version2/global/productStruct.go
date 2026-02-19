package global

type ProductStruct struct {
	Product_id     int     `json:"product_id"`
	Product_name   string  `json:"product_name"`
	Manufacturer   string  `json:"manufacturer"`
	Seller_user_id int     `json:"seller_user_id"`
	Description    string  `json:"description"`
	Screen_size    float32 `json:"screen_size"`
	Picture_url    string  `json:"picture_url"`
	Price          float32 `json:"price"`
	Stock          int     `json:"stock"`
	Active         bool    `json:"active"`
}
