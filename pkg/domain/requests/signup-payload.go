package requests

type SignUpPayload struct {
	Uname    string `json:"uname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
