package msghandler

type Msg struct {
	Data []byte `json:"data"`
}

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (r *Handler) Handle(msg *Msg) error {
	return nil
}
