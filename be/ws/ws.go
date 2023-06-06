package ws

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/blinkinglight/things/fe2"
	"github.com/blinkinglight/things/shared"

	"github.com/gorilla/websocket"
)

type Query struct {
	Subject string `json:"subject"`
	Payload string `json:"payload"`
}

type Component struct {
	Name     string                 `json:"name"`
	Props    map[string]interface{} `json:"props"`
	Function string                 `json:"function,omitempty"`
}

type Row struct {
	CCount     string      `json:"c_count"`
	Components []Component `json:"components"`
}
type Response struct {
	Rows   []Row   `json:"rows"`
	OnLoad []Query `json:"onload,omitempty"`
}

func Run(ctx shared.Context) {
	nc := ctx.Nats()
	_ = nc

	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "X-Token, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("content-security-policy", "default-src 'none'")

		log.Println("got request")
		subject := r.URL.Query().Get("subject")
		requestType := r.URL.Query().Get("type")
		me := r.URL.Query().Get("me")
		blocking := r.URL.Query().Get("block")
		_, _, _ = subject, requestType, me

		switch requestType {
		case "command":
			b, err := ioutil.ReadAll(r.Body)
			log.Printf("wwww %v", err)
			log.Printf("aaa %s %v", string(b), subject)
			var request shared.Message
			err = request.Unmarshal(string(b))
			log.Printf("vvv %v %+v", err, request)

			var payload shared.Message
			payload.SetMetadata("respond", false)
			payload.SetMetadata("request", request.Data)
			payload.SetMetadata("reply_to", "abra")
			payload.SetMetadata("place", request.GetMetadata("place"))
			data, _ := payload.Marshal()
			if blocking == "1" {
				msg, err := nc.Conn().Request(subject, []byte(data), 5*time.Second)
				if err != nil {
					log.Printf("error: %v", err)
					w.Write([]byte(`{"success": 0}`))
					return
				}
				w.Write(msg.Data)
				return
			}
			log.Printf("topic: %v", subject)
			err = nc.Conn().Publish(subject, []byte(data))
			if err != nil {
				log.Printf("error: %v", err)
				w.Write([]byte(`{"success": 0}`))
				return
			}
			w.Write([]byte(`{"success": 1}`))
			return
		case "query":
			b, _ := ioutil.ReadAll(r.Body)

			var request shared.Message
			err := request.Unmarshal(string(b))
			log.Printf("v %v", err)

			var payload shared.Message

			payload.SetMetadata("request", request)
			data, _ := payload.Marshal()
			msg, err := nc.Conn().Request(subject, []byte(data), 5*time.Second)
			if err != nil {
				log.Printf("error: %v", err)
				w.Write([]byte(`{"success": 0}`))
				return
			}
			w.Write(msg.Data)
		}

	}

	upgrader := websocket.Upgrader{}
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	// Define the WebSocket handler
	wshandler := func(w http.ResponseWriter, r *http.Request) {
		// Upgrade the HTTP connection to WebSocket
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("WebSocket upgrade failed:", err)
			return
		}
		defer conn.Close()

		// Subscribe to NATS subject and forward messages to WebSocket
		sub, err := nc.Conn().SubscribeSync("abra")
		if err != nil {
			log.Println("NATS subscription failed:", err)
			return
		}
		defer sub.Unsubscribe()

		// Forward messages from NATS to WebSocket
		go func() {
			for {
				msg, err := sub.NextMsg(time.Hour)
				if err != nil {
					log.Println("NATS message subscription failed:", err)
					conn.Close()
					return
				}

				err = conn.WriteMessage(websocket.TextMessage, msg.Data)
				if err != nil {
					log.Println("WebSocket write message failed:", err)
					return
				}
			}
		}()

		// Forward messages from WebSocket to NATS
		for {
			_, data, err := conn.ReadMessage()
			if err != nil {
				log.Println("WebSocket read message failed:", err)
				return
			}

			err = nc.Publish("abra", data)
			if err != nil {
				log.Println("NATS message publish failed:", err)
				return
			}
		}
	}

	http.HandleFunc("/fe", func(w http.ResponseWriter, r *http.Request) {

		var buff bytes.Buffer
		c := fe2.Posts([]fe2.Post{
			{
				Name:   "Test",
				Author: "autor 1",
			},
		})
		c.Render(context.Background(), &buff)
		w.Write(buff.Bytes())
	})

	http.HandleFunc("/pipe", handler)
	http.HandleFunc("/ws", wshandler)
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		// log.Println("got request")
		response := &Response{
			OnLoad: []Query{
				{
					Subject: "svc.store",
					Payload: "fe.rerender-table",
				},
			},
			Rows: []Row{
				{
					CCount: "col-4",
					Components: []Component{
						{
							Name: "Test",
							Props: map[string]interface{}{
								"color": "red",
								"size":  "large",
							},
						},
						{
							Name: "Test",
							Props: map[string]interface{}{
								"color": "red1",
								"size":  "large1",
							},
						},
					},
				},
				{
					CCount: "col-12",
					Components: []Component{
						{
							Name: "InputV1",
							Props: map[string]interface{}{
								"command": "svc.store",
							},
						},
					},
				},
				{
					CCount: "col-12",
					Components: []Component{
						{
							Name: "TableV1",
							Props: map[string]interface{}{
								"command": "svc.store",
							},
							Function: "fe.rerender-table",
						},
					},
				},
			},
		}
		b, _ := json.Marshal(response)
		w.Write(b)
	})
	http.HandleFunc("/api/update", func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte(`{"success": 0}`))
			return
		}
		log.Printf("%s", b)

		w.Write([]byte(`{"success": 1}`))
	})
	http.HandleFunc("/", http.FileServer(http.Dir(os.Getenv("HTTP_ROOT"))).ServeHTTP)
	log.Println("HTTP is listening on :3000")
	// log.Fatal(http.ListenAndServe(":3000", nil))
	server := http.Server{
		Addr: ":3000",
	}
	go func() {
		log.Fatal(server.ListenAndServe())
	}()
	<-ctx.Ctx.Done()
	server.Shutdown(context.TODO())
}
