package shared

import "encoding/json"

type Structure map[string]interface{}

func (s Structure) Set(key string, value interface{}) {
	s[key] = value
}

func (s Structure) Get(key string) (interface{}, bool) {
	value, ok := s[key]
	if ok {
		return value, ok
	}
	return nil, ok
}

func (s Structure) GetStructure(key string) (Structure, bool) {
	value, ok := s[key]
	if !ok {
		return nil, ok
	}
	st := Structure{}
	for k, v := range value.(map[string]interface{}) {
		(st)[k] = v
	}
	return st, ok
}

func (s Structure) GetString(key string) (string, bool) {
	value, ok := s[key]
	switch value.(type) {
	case string:
		return value.(string), ok
	default:
		return "", ok
	}
}

type Message struct {
	Data     Structure `json:"data"`
	Metadata Structure `json:"metadata"`
}

func (m *Message) Context() Structure {
	value, ok := m.Metadata.Get("context")
	if !ok {
		return nil
	}
	return value.(Structure)
}

func (m *Message) CopyContextFrom(msg Message) {
	context := msg.Context()
	if context != nil {
		m.Metadata.Set("context", context)
	}
}

func (m *Message) Request(key string) (string, bool) {
	value, ok := m.Metadata.GetStructure("request")
	if !ok {
		return "", ok
	}
	v, ok := value.GetString(key)
	return v, ok
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
