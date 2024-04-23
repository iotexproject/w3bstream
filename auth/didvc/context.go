package didvc

import "context"

type clientIDCtxKey struct{}

func WithClientID(parent context.Context, clientID string) context.Context {
	if parent == nil {
		panic("with client id context, nil context")
	}
	return context.WithValue(parent, clientIDCtxKey{}, clientID)
}

func ClientIDFrom(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(clientIDCtxKey{}).(string)
	return v, ok
}
