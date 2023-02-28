package handlers

var pool = make(map[string]IHandler)

type HandlerItem map[string]IHandler

func Register(items ...HandlerItem) {
	for _, item := range items {
		for k, v := range item {
			pool[k] = v
		}
	}
}

func Get(key string) (IHandler, bool) {
	if h, ok := pool[key]; ok {
		return h, true
	}

	return nil, false
}
