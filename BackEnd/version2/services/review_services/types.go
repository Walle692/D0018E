package review_services

type WriteReviewRequest struct {
	Product_id int    `json:"product_id"`
	Comment    string `json:"comment"`
	Rating     int    `json:"rating"`
}

type DeleteReviewRequest struct {
	Comment_id int `json:"comment_id"`
}
