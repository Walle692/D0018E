package products_services

type CreateProductRequest struct {
	ProductName  string  `json:"product_name" binding:"required"`
	Manufacturer string  `json:"manufacturer" binding:"required"`
	Description  string  `json:"description" binding:"required"`
	ScreenSize   int     `json:"screen_size" binding:"required"`
	PictureURL   string  `json:"picture_url" binding:"required"`
	Price        float64 `json:"price" binding:"required"`
	Stock        int     `json:"stock" binding:"required"`
}

type DeleteProduct struct {
	Product_id int `json:"product_id"`
}
