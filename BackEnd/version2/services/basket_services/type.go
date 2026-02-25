package basket_services

// takes the product id and quantity from the request and adds the product to the user's basket
type AddToBasketRequest struct {
	ProductID int `json:"product_id" binding:"required"`
	Quantity  int `json:"quantity" binding:"required,min=1"`
}

type DeleteBasketRequest struct {
	ProductID int `json:"product_id" binding:"required"`
}
