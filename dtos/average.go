package dtos

type AverageMoneyPerDayResponse struct {
	Money int64 `json:"money_spend_per_day"`
}

type AvrMoneyMonthRequest struct {
	UserID string
	Month  int64 `validate:"max=12,min=1"`
}
