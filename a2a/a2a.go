package a2a

import (
	"github.com/franzego/stage03/models"
)

func Responder(event, user, message string) *models.TelexResponse {
	return &models.TelexResponse{
		EventName: event,
		Username:  user,
		Message:   message,
	}
}
