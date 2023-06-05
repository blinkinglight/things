// Code generated by templ@(devel) DO NOT EDIT.

package card

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func cb(payload , topic string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_cb_9232`,
		Function: `function __templ_cb_9232(payload, topic){console.log(payload);
	let url = "http://localhost:3000/pipe?type=command&me=abra&subject="+topic;
	  return fetch(url, {
          method: "POST",
          body: JSON.stringify({data: {id : payload}}),
          headers: {
              "Content-Type": "application/json"
          }
      }).then(res => res.json()).then(res => {
		  console.log(res);
	  }).catch(err => {
		  console.log(err);
	  });}`,
		Call: templ.SafeScript(`__templ_cb_9232`, payload, topic),
	}
}

func cardTemplate(card Card) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		// Element (standard)
		_, err = templBuffer.WriteString("<div")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" class=\"card\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<div")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" class=\"card-header\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<h3")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" class=\"card-title\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// StringExpression
		var var_2 string = card.Title
		_, err = templBuffer.WriteString(templ.EscapeString(var_2))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h3>")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<div")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" class=\"card-tools\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<span")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" class=\"badge badge-primary\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Text
		var_3 := `Label`
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</span>")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<div")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" class=\"card-body\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// CallTemplate
		err = card.Body.Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func Widget(card Card) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_4 := templ.GetChildren(ctx)
		if var_4 == nil {
			var_4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		// TemplElement
		err = cardTemplate(card).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}
