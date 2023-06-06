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

	for _, s := range shared.Registry {
		s := s
		go func() {
			ctx, err := shared.NewContext()
			if err != nil {
				panic(err)
			}
			srv, _ := service.AddService(ctx, s)
			log.Printf("service started: %s", srv.Info().Name)
		}()
	}
	ctx, err := shared.NewContext()
	if err != nil {
		panic(err)
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
