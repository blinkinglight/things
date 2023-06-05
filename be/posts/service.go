package posts

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/blinkinglight/things/fe2"
	"github.com/blinkinglight/things/shared"
)

func init() {
	shared.RegisterService(shared.Service{
		Subject: "svc.posts",
		Fn:      Run,
		Name:    "PostsService",
		Version: "1.0.0",
	})
	shared.RegisterService(shared.Service{
		Subject: "svc.post.delete",
		Fn:      RunDelete,
		Name:    "PostsDeleteService",
		Version: "1.0.0",
	})
}

var (
	posts = []fe2.Post{}
)

func Run(ctx shared.Context, message shared.Message) (shared.Message, error) {
	log.Printf("svc.posts got command")

	var msg shared.Message
	var buff bytes.Buffer

	posts = append(posts, fe2.Post{
		ID:     fmt.Sprintf("%d", rand.Intn(1000)),
		Name:   fmt.Sprintf("Test %d", len(posts)+1),
		Author: fmt.Sprintf("autor %d", len(posts)+1),
	})

	fe2.Posts(posts).Render(context.Background(), &buff)
	msg.SetData("html", buff.String())
	msg.SetData("place", "html")

	return msg, nil
}

func RunDelete(ctx shared.Context, message shared.Message) (shared.Message, error) {
	log.Printf("svc.post.delete got c	ommand %v", message)

	var msg shared.Message
	var buff bytes.Buffer
	var tmp []fe2.Post
	for _, post := range posts {
		// tmp = append(tmp, post)
		log.Printf("%v == %v", message.GetMetadata("request").(map[string]interface{})["id"].(string), post.ID)
		if post.ID != message.GetMetadata("request").(map[string]interface{})["id"].(string) {
			tmp = append(tmp, post)
		}
	}
	posts = tmp
	fe2.Posts(tmp).Render(context.Background(), &buff)
	msg.SetData("html", buff.String())
	msg.SetData("place", "html")

	return msg, nil
}
