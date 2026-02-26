package review

type Review struct {
	Comment_id       int    `json:"comment_id"`
	Customer_user_id int    `json:"customer_user_id"`
	Product_id       int    `json:"product_id"`
	Comment          string `json:"comment"`
	Rating           int    `json:"rating"`
}
