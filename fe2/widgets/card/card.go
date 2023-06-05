package card

import "github.com/a-h/templ"

type Card struct {
	Title string          `json:"title"`
	Body  templ.Component `json:"body"`
}
