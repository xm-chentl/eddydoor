package ctxtype

type Value string

func (v Value) String() string {
	return string(v)
}

const (
	Err     Value = "ctx_error"
	Route   Value = "ctx_route"
	Handler Value = "ctx_handler"
)
