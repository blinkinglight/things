package shared

type Wrapper func(Context, Message) (Message, error)

func (w Wrapper) Execute(ctx Context, m Message) (Message, error) {
	return w(ctx, m)
}
