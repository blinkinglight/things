package user

import (
	"log"

	"github.com/blinkinglight/thingsbe/shared"
)

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	SessionID string `json:"session_id"`
	Success   int    `json:"success"`
	Message   string `json:"message,omitempty"`
}

func init() {
	shared.RegisterService(shared.Service{
		Subject: "svc.user",
		Fn:      Run,
		Name:    "UserService",
		Version: "1.0.0",
	})
}

func Run(ctx shared.Context, message shared.Message) (shared.Message, error) {
	log.Printf("svc.login got command")

	log.Printf("%+v", message)

	var msg shared.Message
	msg.Data = map[string]interface{}{
		"success": 1,
		"message": "user in",
		"@type":   "Response",
	}

	return msg, nil
}
