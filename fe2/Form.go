package fe2

type Form struct {
	ID      string `json:"data-id"`
	Command string `json:"data-function"`
	Target  string `json:"data-target"`
}
