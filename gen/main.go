package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var packageTemplate = template.Must(template.New("").Parse(`package {{ .LName }}

import (
	"github.com/blinkinglight/things/shared"
	"log"
)

func init() {
	shared.RegisterService(shared.Service{
		Subject: "{{ .Subject}}",
		Fn:      Run,
		Name:    "{{ .Name }}Service",
		Version: "1.0.0",
	})
}

func Run(ctx shared.Context, message shared.Message) (shared.Message, error) {
	log.Printf("{{ .Subject }} got command")

	var msg shared.Message

	return msg, nil
}

`))

func main() {
	if len(os.Args) < 4 {
		log.Fatal("usage: gen <folder name> <service name> <service subject>")
	}
	folderName := os.Args[1]
	serviceName := os.Args[2]
	serviceSubject := os.Args[3]
	os.MkdirAll("./be/"+folderName, 0755)

	fh, err := os.Create("./be/" + folderName + "/service.go")
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()

	packageTemplate.Execute(fh, struct {
		Name    string
		LName   string
		Subject string
	}{
		Name:    serviceName,
		LName:   strings.ToLower(serviceName),
		Subject: serviceSubject,
	})
}
