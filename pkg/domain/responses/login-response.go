package responses

import "github.com/ong-gtp/go-chat/pkg/models"

type LoginResponse struct {
	User     models.User `json:"User"`
	JwtToken string      `json:"Token"`
}
