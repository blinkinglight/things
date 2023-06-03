package main

import (
	"log"

	"github.com/9glt/go-signals"
	"github.com/blinkinglight/things/be/ws"
	"github.com/blinkinglight/things/shared"

	"github.com/nats-io/nats.go/micro"
)

func main() {

	ctx, err := shared.NewContext()
	if err != nil {
		panic(err)
	}

	for _, s := range shared.Registry {
		s := s
		go func() {
			srv, err := micro.AddService(ctx.Nats().Conn(), micro.Config{
				Name:    s.Name,
				Version: s.Version,
				// base handler
				Endpoint: &micro.EndpointConfig{
					Subject: s.Subject,
					Handler: micro.HandlerFunc(func(req micro.Request) {
						log.Printf("%s got command", s.Name)
						message := shared.Message{}
						err := message.Unmarshal(string(req.Data()))
						if err != nil {
							req.Error("503", "Internal error", []byte(err.Error()))
							return
						}

						response, err := s.Fn.Execute(ctx, message)
						if err != nil {
							req.Error("503", "Internal error", []byte(err.Error()))
							return
						}
						responseData, err := response.Marshal()
						if err != nil {
							req.Error("503", "Internal error", []byte(err.Error()))
							return
						}

						if message.GetMetadata("internal") != "true" {
							req.Respond([]byte(`{"status":"ok"}`))
						} else {
							req.Respond([]byte(responseData))
						}

						if message.GetMetadata("reply_to") != nil {
							ctx.Nats().Conn().Publish(message.GetMetadata("reply_to").(string), []byte(responseData))
						}
					}),
				},
			})
			if err != nil {
				panic(err)
			}
			defer srv.Stop()
			log.Printf("service started: %s", srv.Info().Name)

			for {
				select {
				case <-ctx.Ctx.Done():
					return
				}
			}
		}()
	}

	go ws.Run(ctx)

	go func() {
		signals.INT(func() {
			log.Println("got interrupt")
			ctx.Cancel()
		})
	}()
	ctx.Wait()

}
