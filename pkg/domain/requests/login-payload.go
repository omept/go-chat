package requests

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
