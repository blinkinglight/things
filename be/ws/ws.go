package ws

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/blinkinglight/things/shared"

	"github.com/gorilla/websocket"
)

func Run(ctx shared.Context) {
	nc := ctx.Nats()
	_ = nc

	handler := func(w http.ResponseWriter, r *http.Request) {
		log.Println("got request")
		subject := r.URL.Query().Get("subject")
		requestType := r.URL.Query().Get("type")
		me := r.URL.Query().Get("me")
		_, _, _ = subject, requestType, me

		switch requestType {
		case "command":
			b, _ := ioutil.ReadAll(r.Body)
			var request shared.Message
			err := request.Unmarshal(string(b))
			log.Printf("%v", err)

			var payload shared.Message
			payload.SetMetadata("respond", false)
			payload.SetMetadata("request", request.Data)
			payload.SetMetadata("reply_to", "abra")
			data, _ := payload.Marshal()
			msg, err := nc.Conn().Request(subject, []byte(data), 5*time.Second)
			if err != nil {
				log.Printf("error: %v", err)
				return
			}
			w.Write(msg.Data)
		case "query":
			b, _ := ioutil.ReadAll(r.Body)

			var request shared.Message
			err := request.Unmarshal(string(b))
			log.Printf("%v", err)

			var payload shared.Message

			payload.SetMetadata("request", request)
			data, _ := payload.Marshal()
			msg, err := nc.Conn().Request(subject, []byte(data), 5*time.Second)
			if err != nil {
				log.Printf("error: %v", err)
				return
			}
			w.Write(msg.Data)
		}

	}

	upgrader := websocket.Upgrader{}

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

	http.HandleFunc("/pipe", handler)
	http.HandleFunc("/ws", wshandler)
	http.HandleFunc("/", http.FileServer(http.Dir(os.Getenv("HTTP_ROOT"))).ServeHTTP)
	log.Println("HTTP is listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
