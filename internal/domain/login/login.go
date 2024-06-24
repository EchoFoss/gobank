package domain

type LoginRequest struct {
	Number   int    `json:"number"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Number int    `json:"number"`
	Token  string `json:"token"`
}
