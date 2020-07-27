package dto

type UserDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`

	Token    string `json:"token"`
	Auth     string `json:"auth"`
}
