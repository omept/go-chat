package responses

import "github.com/ong-gtp/go-chat/models"

type SignUpResponse struct {
	User     models.User `json:"User"`
	JwtToken string      `json:"Token"`
}
