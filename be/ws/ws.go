package ws

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Message struct {
	UserID    string          `json:"userID,omitempty"`
	SessionID string          `json:"sessionID,omitempty"`
	Type      string          `json:"type"`
	Name      string          `json:"name"`
	ReplyTo   string          `json:"replyTo"`
	Payload   json.RawMessage `json:"payload"`
}

func Run(ctx context.Context) {
	// Set up the HTTP server
	http.HandleFunc("/", http.FileServer(http.Dir(os.Getenv("HTTP_ROOT"))).ServeHTTP)
	log.Println("HTTP is listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
