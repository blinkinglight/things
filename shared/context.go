package shared

import (
	"context"
	"time"
)

type Context struct {
	Ctx context.Context
	cfn context.CancelFunc
	nc  *NATS
}

func NewContext() (Context, error) {
	var ctx Context
	ctx.Ctx, ctx.cfn = context.WithCancel(context.Background())
	var err error
	ctx.nc, err = NewNATS()
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}

func (ctx Context) Cancel() {
	ctx.cfn()
}

func (ctx Context) Close() {
	ctx.nc.Close()
}

func (ctx Context) Wait() {
	<-ctx.Ctx.Done()
}

func (ctx Context) Nats() *NATS {
	return ctx.nc
}

func (ctx Context) Request(subject string, message Message) (Message, error) {
	var response Message
	message.SetMetadata("internal", "true")
	data, err := message.Marshal()
	if err != nil {
		return response, err
	}
	msg, err := ctx.nc.Conn().Request(subject, []byte(data), 5*time.Second)
	if err != nil {
		return response, err
	}
	err = response.Unmarshal(string(msg.Data))
	if err != nil {
		return response, err
	}
	return response, nil
}

func (ctx Context) Publish(subject string, message Message) error {
	data, err := message.Marshal()
	if err != nil {
		return err
	}
	err = ctx.nc.Conn().Publish(subject, []byte(data))
	if err != nil {
		return err
	}
	return nil
}
