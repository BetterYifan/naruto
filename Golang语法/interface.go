package Golang语法

type Handle interface {
	Add(i int64)
}

type HttpHandler struct {
}

func (h *HttpHandler) Add(i int64) {
	return
}

var _ Handle = (*HttpHandler)(nil)
