package main

import (
	"log"
	"time"

	signals "github.com/9glt/go-signals"
	"github.com/blinkinglight/things/be/ws"
	"github.com/blinkinglight/things/service"
	"github.com/blinkinglight/things/shared"
)

func main() {

	ctx, err := shared.NewContext()
	if err != nil {
		panic(err)
	}

	for _, s := range shared.Registry {
		s := s
		go func() {
			srv, _ := service.AddService(ctx, s)
			log.Printf("service started: %s", srv.Info().Name)
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
	time.Sleep(1 * time.Second)
}
