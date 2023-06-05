package alive

import (
	"bytes"
	"context"
	"log"

	"github.com/blinkinglight/things/fe2"
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
	// for i := 0; i < 10; i++ {
	// 	m := shared.Message{}
	// 	m.SetData("ping", fmt.Sprintf("pong %d", i))
	// 	ctx.Publish(message.GetMetadata("reply_to").(string), m)
	// 	time.Sleep(1 * time.Second)
	// }

	var msg shared.Message

	var buff bytes.Buffer
	fe2.Home().Render(context.Background(), &buff)
	// msg.SetData("html", buff.String())
	ctx.Publish(message.GetMetadata("reply_to").(string), shared.Message{
		Data: map[string]interface{}{
			"html": buff.String(),
		},
	})

	return msg, nil
}
