package clients

import "context"

type clientIDCtxKey struct{}

func WithClientID(parent context.Context, client *Client) context.Context {
	if parent == nil {
		panic("with client id context, nil context")
	}
	return context.WithValue(parent, clientIDCtxKey{}, client)
}

func ClientIDFrom(ctx context.Context) *Client {
	v, _ := ctx.Value(clientIDCtxKey{}).(*Client)
	return v
}
