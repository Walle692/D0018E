package basket

type BasketResponse struct {
	Totalprice   float64              `json:"totalprice"`
	Basket_items []BasketItemResponse `json:"basket_items"`
}

type BasketItemResponse struct {
	Product_id   int     `json:"product_id"`
	Product_name string  `json:"product_name"`
	Manufacturer string  `json:"manufacturer"`
	Picture_url  string  `json:"picture_url"`
	Price        float64 `json:"price"`
	Quantity     int     `json:"quantity"`
	Available    bool    `json:"available"`
}
