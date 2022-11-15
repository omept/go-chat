package requests

type SignUpPayload struct {
	Uname    string `json:"uname"`
	Password string `json:"password"`
}
