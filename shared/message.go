package shared

import "encoding/json"

type Message struct {
	Data     map[string]interface{} `json:"data"`
	Metadata map[string]interface{} `json:"metadata"`
}

func (m *Message) Error() string {
	b, _ := json.Marshal(m)
	return string(b)
}

func (m *Message) Marshal() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

func (m *Message) Unmarshal(b string) error {
	return json.Unmarshal([]byte(b), m)
}

func (m *Message) SetData(key string, value interface{}) {
	if m.Data == nil {
		m.Data = make(map[string]interface{})
	}
	m.Data[key] = value
}

func (m *Message) CopyMetadataFrom(msg Message) {
	m.Metadata = msg.Metadata
}

func (m *Message) SetMetadata(key string, value interface{}) {
	if m.Metadata == nil {
		m.Metadata = make(map[string]interface{})
	}
	m.Metadata[key] = value
}

func (m *Message) GetMetadata(key string) interface{} {
	if m.Metadata == nil {
		return ""
	}
	return m.Metadata[key]
}

func (m *Message) GetData(key string) interface{} {
	return m.Data[key]
}

func (m *Message) GetSubject() string {
	return m.GetMetadata("subject").(string)
}

func (m *Message) SetSubject(subject string) {
	m.SetMetadata("subject", subject)
}

func (m *Message) GetReplyTo() string {
	return m.GetMetadata("replyTo").(string)
}

func (m *Message) SetReplyTo(replyTo string) {
	m.SetMetadata("replyTo", replyTo)
}
