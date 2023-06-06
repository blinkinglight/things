package service

import (
	"log"
	"sync"

	"github.com/blinkinglight/things/shared"
	"github.com/nats-io/nats.go/micro"
)

func AddService(ctx shared.Context, s shared.Service) (micro.Service, error) {
	var err error
	var svc micro.Service
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		svc, err = micro.AddService(ctx.Nats().Conn(), micro.Config{
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
						log.Printf("%v", err)
						req.Error("503", "Internal error", []byte(err.Error()))
						return
					}
					// ctx.SetMessage(message)
					log.Printf("message: %+v", message)
					c := ctx.Clone()
					c.SetMessage(message)
					response, err := s.Fn.Execute(c, message)
					if err != nil {
						req.Error("503", "Internal error", []byte(err.Error()))
						c.Publish("abra", shared.Message{
							Data: map[string]interface{}{
								"error": err.Error(),
							},
						})
						return
					}
					log.Printf("place here %+v", message.GetMetadata("place"))
					if message.GetMetadata("place") != nil {
						response.SetMetadata("place", message.GetMetadata("place"))
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
						c.Nats().Conn().Publish(message.GetMetadata("reply_to").(string), []byte(responseData))
					}
				}),
			},
		})
		if err != nil {
			panic(err)
		}
		defer svc.Stop()
		wg.Done()
		for {
			select {
			case <-ctx.Ctx.Done():
				return
			}
		}
	}()
	wg.Wait()
	return svc, err
}
