package requests

type LoginPayload struct {
	Uname    string `json:"uname"`
	Password string `json:"password"`
}
