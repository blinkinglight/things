package popup

import (
	"log"

	"github.com/blinkinglight/things/shared"
)

func init() {
	shared.RegisterService(shared.Service{
		Subject: "svc.popup",
		Fn:      Run,
		Name:    "PopupService",
		Version: "1.0.0",
	})
}

func Run(ctx shared.Context, message shared.Message) (shared.Message, error) {
	log.Printf("svc.popup got command")

	var msg shared.Message

	msg.SetData("command", "popup")
	msg.SetData("text", "hello world")
	msg.SetData("title", "hello world")
	msg.SetData("button", map[string]interface{}{
		"command": map[string]interface{}{
			"subject": "svc.popup",
			"type":    "command",
			"action":  "close",
		},
	})

	return msg, nil
}
