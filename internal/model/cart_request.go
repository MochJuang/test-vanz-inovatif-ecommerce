package model

type AddToCartRequest struct {
	UserID    uint `json:"user_id" validate:"required"`
	ProductID uint `json:"product_id" validate:"required"`
	Quantity  int  `json:"quantity" validate:"required"`
}
