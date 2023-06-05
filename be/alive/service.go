package alive

import (
	"fmt"
	"log"
	"time"

	"github.com/blinkinglight/things/shared"
)

func init() {
	shared.RegisterService(shared.Service{
		Subject: "svc.alive",
		Fn:      Run,
		Name:    "AliveService",
		Version: "1.0.0",
	})
}

func Run(ctx shared.Context, message shared.Message) (shared.Message, error) {
	log.Printf("svc.alive got command")
	for i := 0; i < 10; i++ {
		m := shared.Message{}
		m.SetData("ping", fmt.Sprintf("pong %d", i))
		ctx.Publish(message.GetMetadata("reply_to").(string), m)
		time.Sleep(1 * time.Second)
	}

	var msg shared.Message

	return msg, nil
}
