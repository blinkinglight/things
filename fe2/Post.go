package fe2

import "fmt"

type Post struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

func (p *Post) Call(name, id string) string {
	return fmt.Sprintf(`call("%s", "%s");`, name, id)
}
