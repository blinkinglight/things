package shared

import "encoding/json"

type CommandPayload struct {
	Subject string `json:"subject"`
	ReplyTo string `json:"replyTo",omitempty`
	Payload string `json:"payload"`
	Success int    `json:"success"`
}

func (p *CommandPayload) Marshal() ([]byte, error) {
	return json.Marshal(p)
}

type QueryPayload struct {
	Subject string `json:"subject"`
	ReplyTo string `json:"replyTo",omitempty`
	Payload string `json:"payload"`
	Success int    `json:"success"`
}

func (p *QueryPayload) Marshal() ([]byte, error) {
	return json.Marshal(p)
}
