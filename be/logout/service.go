package logout

import (
	"context"
	"cqrs/be/shared"
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go/micro"
)

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Run(ctx context.Context) {
	// TODO
	nc, err := shared.NewNATS()
	if err != nil {
		panic(err)
	}
	_ = nc

	// request handler
	echoHandler := func(req micro.Request) {
		log.Printf("svc.logout got command")
		var auth Auth
		err := json.Unmarshal(req.Data(), &auth)
		if err != nil {
			req.Error("503", "Internal error", []byte(err.Error()))
			return
		}
		req.Respond([]byte(`{"success":1, "message": "logged out"}`))
	}

	srv, err := micro.AddService(nc.Conn(), micro.Config{
		Name:    "LogoutService",
		Version: "1.0.0",
		// base handler
		Endpoint: &micro.EndpointConfig{
			Subject: "svc.logout",
			Handler: micro.HandlerFunc(echoHandler),
		},
	})
	defer srv.Stop()

	log.Printf("service started: %s", srv.Info().Name)

	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}
