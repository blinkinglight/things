package login

import (
	"log"

	"github.com/blinkinglight/things/be/shared"
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
		Subject: "svc.login",
		Fn:      Run,
		Name:    "LoginService",
		Version: "1.0.0",
	})
}

func Run(ctx shared.Context, message shared.Message) (shared.Message, error) {
	log.Printf("svc.login got command")

	m, err := ctx.Request("svc.user", shared.Message{
		Data: map[string]interface{}{
			"username": "admin",
			"password": "admin",
			"@type":    "Auth",
		},
	})
	if err != nil {
		return shared.Message{}, &shared.Message{
			Data: map[string]interface{}{
				"success": 0,
				"message": err.Error(),
				"@type":   "Response",
			},
		}
	}

	log.Printf("%+v", m)

	var msg shared.Message
	msg.Data = map[string]interface{}{
		"session_id": "7d652daa88df0c3fef13e23e99ca9e4ce5e31902",
		"success":    1,
		"message":    "logged in",
		"@type":      "Response",
	}
	msg.Metadata = map[string]interface{}{
		"message": m.Data,
	}

	return msg, nil
}
