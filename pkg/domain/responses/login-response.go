package responses

import "github.com/ong-gtp/go-chat/pkg/models"

type LoginResponse struct {
	User     models.User `json:"user"`
	JwtToken string      `json:"token"`
}
