package model

type PaymentRequest struct {
	Remarks string  `json:"remarks"`
	Amount  float64 `json:"amount"`
}
