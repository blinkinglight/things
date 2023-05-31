package main

import (
	"context"
	"cqrs/be/login"
	"cqrs/be/logout"
	"cqrs/be/ws"
)

func main() {

	ctx := context.Background()

	go login.Run(ctx)
	go logout.Run(ctx)
	go ws.Run(ctx)

	select {}
}
