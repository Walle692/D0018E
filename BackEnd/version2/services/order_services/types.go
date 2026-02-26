package order_services

type OrderSearchRequest struct {
	SearchType   string `json:"search_type" binding:"required"`
	SearchObject string `json:"search_object" binding:"required"`
}
