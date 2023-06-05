package store

import (
	"log"

	"github.com/blinkinglight/things/shared"
)

func init() {
	shared.RegisterService(shared.Service{
		Subject: "svc.store",
		Fn:      Run,
		Name:    "StoreService",
		Version: "1.0.0",
	})
}

type Item struct {
	Type  string                 `json:"type"`
	Value string                 `json:"value"`
	Props map[string]interface{} `json:"props"`
}

type Response struct {
	Fields []Item `json:"fields"`
}

var (
	items = []Item{}
)

func Run(ctx shared.Context, message shared.Message) (shared.Message, error) {
	log.Printf("svc.store got command %+v", message)

	if message.GetMetadata("request") != nil {
		items = append(items, Item{
			Type:  "TableRowV1",
			Value: message.GetMetadata("request").(map[string]interface{})["todo"].(string),
			Props: map[string]interface{}{
				"value": message.GetMetadata("request").(map[string]interface{})["todo"].(string),
			},
		})
	}

	var msg shared.Message
	msg.SetMetadata("command", "fe.rerender-table")
	msg.SetData("data", map[string]interface{}{
		"table": "store",
		"rows":  Response{Fields: items},
	})

	return msg, nil
}
