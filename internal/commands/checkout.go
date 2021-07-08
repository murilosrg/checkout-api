package commands

type Checkout struct {
	TotalAmount             int `json:"total_amount"`
	TotalAmountWithDiscount int `json:"total_amount_with_discount"`
	TotalDiscount           int `json:"total_discount"`
	Products                []ProductResponse `json:"products"`
}

type ProductResponse struct {
	ID          int  `json:"id"`
	Quantity    int  `json:"quantity"`
	UnitAmount  int  `json:"unit_amount"`
	TotalAmount int  `json:"total_amount"`
	Discount    int  `json:"discount"`
	IsGift      bool `json:"is_gift"`
}

func NewProductGift(id int) ProductResponse {
	return ProductResponse{ID: id, Quantity: 1, IsGift: true}
}
