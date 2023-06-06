package posts

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"

	"github.com/a-h/templ"
	"github.com/blinkinglight/things/fe2"
	"github.com/blinkinglight/things/fe2/widgets/card"
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
		Subject: "svc.post.render",
		Fn:      RunRender,
		Name:    "PostsRenderService",
		Version: "1.0.0",
	})
	shared.RegisterService(shared.Service{
		Subject: "svc.post.edit",
		Fn:      RunEdit,
		Name:    "PostsEditService",
		Version: "1.0.0",
	})
	shared.RegisterService(shared.Service{
		Subject: "svc.post.store",
		Fn:      RunStore,
		Name:    "PostsStoreService",
		Version: "1.0.0",
	})
	shared.RegisterService(shared.Service{
		Subject: "svc.post.delete",
		Fn:      RunDelete,
		Name:    "PostsDeleteService",
		Version: "1.0.0",
	})
	shared.RegisterService(shared.Service{
		Subject: "svc.post.create",
		Fn:      RunCreate,
		Name:    "PostsCreateService",
		Version: "1.0.0",
	})
}

var (
	posts = []fe2.Post{}
)

func Unsafe(html string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, html)
		return
	})
}

func Run(ctx shared.Context, message shared.Message) (shared.Message, error) {
	log.Printf("svc.posts got command")

	var msg shared.Message

	posts = append(posts, fe2.Post{
		ID:     fmt.Sprintf("%d", rand.Intn(1000)),
		Name:   fmt.Sprintf("Test %d", len(posts)+1),
		Author: fmt.Sprintf("autor %d", len(posts)+1),
	})

	ctx.Publish("svc.post.render", message)

	return msg, nil
}

func RunEdit(ctx shared.Context, message shared.Message) (shared.Message, error) {
	log.Printf("svc.post.edit got command %+v", message)

	var msg shared.Message

	var buff bytes.Buffer

	var post fe2.Post
	reqID, ok := message.Request("id")
	if ok {
		for _, p := range posts {
			if p.ID == reqID {
				post = p
				break
			}
		}
	}

	topic := "svc.post.create"
	if ok && reqID != "" {
		topic = "svc.post.store"
	}
	err := fe2.Edit(fe2.Post{
		ID:     post.ID,
		Name:   post.Name,
		Author: post.Author,
	}, fe2.Form{
		ID:      reqID,
		Target:  "html",
		Command: topic,
	},
		`{"shit":"value"}`).Render(context.Background(), &buff)
	log.Printf("%v %s", err, buff.String())

	var badge bytes.Buffer
	card.Badge().Render(context.Background(), &badge)

	var html bytes.Buffer
	card.Widget(card.Card{
		Title: "Edit post",
		Body:  Unsafe(buff.String()),
		Badge: Unsafe(badge.String()),
	}).Render(context.Background(), &html)

	msg.SetData("html", html.String())

	return msg, nil
}

func RunCreate(ctx shared.Context, message shared.Message) (shared.Message, error) {
	log.Printf("svc.posts got command %+v", message)

	var msg shared.Message
	name, _ := message.Request("name")
	author, _ := message.Request("author")
	posts = append(posts, fe2.Post{
		ID:     fmt.Sprintf("%d", rand.Intn(1000)),
		Name:   name,
		Author: author,
	})

	ctx.Publish("svc.post.render", message)

	return msg, nil
}

func RunStore(ctx shared.Context, message shared.Message) (shared.Message, error) {
	log.Printf("svc.post.store got command %+v", message)

	var msg shared.Message

	reqID, ok := message.Request("id")
	if ok {
		for idx, p := range posts {
			if p.ID == reqID {
				posts[idx].Author, _ = message.Request("author")
				posts[idx].Name, _ = message.Request("name")
				break
			}
		}
	}

	ctx.Publish("svc.post.render", message)

	return msg, nil
}

func RunRender(ctx shared.Context, message shared.Message) (shared.Message, error) {
	log.Printf("svc.post.render got command")

	var msg shared.Message
	var buff bytes.Buffer

	var badge bytes.Buffer
	card.Badge().Render(context.Background(), &badge)

	fe2.Posts(posts).Render(context.Background(), &buff)

	var html bytes.Buffer
	card.Widget(card.Card{
		Title: "Posts",
		Body:  Unsafe(buff.String()),
		Badge: Unsafe(badge.String()),
	}).Render(context.Background(), &html)

	msg.SetData("html", html.String())
	// msg.SetData("place", "html")

	return msg, nil
}

func RunDelete(ctx shared.Context, message shared.Message) (shared.Message, error) {
	log.Printf("svc.post.delete got ommand %v", message)

	var msg shared.Message
	// var buff bytes.Buffer
	var tmp []fe2.Post
	id, ok := message.Request("id")
	if !ok {
		log.Printf("debug: %+v", message)
	}
	for _, post := range posts {
		if post.ID != id {
			tmp = append(tmp, post)
		}
	}
	posts = tmp

	ctx.Publish("svc.post.render", message)

	return msg, nil
}
