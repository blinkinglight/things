package logout

import (
	"log"

	"github.com/blinkinglight/things/shared"
)

func init() {
	shared.RegisterService(shared.Service{
		Subject: "svc.logout",
		Fn:      Run,
		Name:    "LogoutService",
		Version: "1.0.0",
	})
}

func Run(ctx shared.Context, message shared.Message) (shared.Message, error) {
	log.Printf("svc.logout got command")

	var msg shared.Message
	msg.Data = map[string]interface{}{
		"success": 1,
		"message": "logged out",
		"@type":   "Response",
	}

	return msg, nil
}
