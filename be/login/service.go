package login

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

type Response struct {
	SessionID string `json:"session_id"`
	Success   int    `json:"success"`
	Message   string `json:"message,omitempty"`
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
		log.Printf("svc.login got command")
		var auth Auth
		err := json.Unmarshal(req.Data(), &auth)
		if err != nil {
			req.Error("503", "Internal error", []byte(err.Error()))
			return
		}
		var resp Response
		resp.SessionID = "7d652daa88df0c3fef13e23e99ca9e4ce5e31902"
		resp.Success = 1
		resp.Message = "logged in"
		b, _ := json.Marshal(resp)
		req.Respond(b)
	}

	srv, err := micro.AddService(nc.Conn(), micro.Config{
		Name:    "LoginService",
		Version: "1.0.0",
		// base handler
		Endpoint: &micro.EndpointConfig{
			Subject: "svc.login",
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
