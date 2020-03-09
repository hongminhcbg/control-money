package dtos

type AvrMoneyResponse struct {
	Average int64 `json:"average"`
}
type AvrMoneyRequest struct {
	UserID string
	Month  int64 `validate:"max=12,min=1"`
}
type Response struct {
	Data interface{} `json:"data"`
	Meta struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"meta"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserID   int64  `json:"id"`
	Username string `json:"username"`
	Name     string `json:"full_name"`
	Tocken   string `json:"tocken"`
	Money    int64  `json:"money"`
}
type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
