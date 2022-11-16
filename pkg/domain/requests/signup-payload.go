package requests

type SignUpPayload struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
