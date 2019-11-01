package web

type userData struct {
	Username             string `json:"username"`
	Email                string `json:"email"`
	PhoneNumber          string `json:"phone_number"`
	Name                 string `json:"name"`
	Role                 string `json:"role"`
	Token                string `json:"token"`
}
