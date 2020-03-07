package dtos

type AvrMoneyResponse struct {
	Average int64 `json:"average"`
}
type AvrMoneyRequest struct {
	UserID string
	Month  int64 `validate:"max=12,min=1"`
}
