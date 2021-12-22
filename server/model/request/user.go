package request

type LoginForm struct {
	Name     string `json:"name" from:"name"`
	Password string `json:"password" from:"password"`
}
