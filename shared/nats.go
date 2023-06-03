package shared

import (
	"os"

	nats "github.com/nats-io/nats.go"
)

type MsgHandler func(msg *nats.Msg)

type NATS struct {
	nc *nats.Conn
}

func NewNATS() (*NATS, error) {
	nc, err := nats.Connect(os.Getenv("NATS_URL"))
	if err != nil {
		return nil, err
	}
	return &NATS{nc: nc}, nil
}

func (n *NATS) Conn() *nats.Conn {
	return n.nc
}

func (n *NATS) Close() {
	n.nc.Close()
}

func (n *NATS) Publish(subject string, data []byte) error {
	return n.nc.Publish(subject, data)
}

func (n *NATS) Subscribe(subject string, cb nats.MsgHandler) (*nats.Subscription, error) {
	return n.nc.Subscribe(subject, cb)
}
