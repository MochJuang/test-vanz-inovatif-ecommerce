package model

type TransferRequest struct {
	TargetUserID string  `json:"target_user_id"`
	Remarks      string  `json:"remarks"`
	Amount       float64 `json:"amount"`
}
