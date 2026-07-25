package decoHandler

import "net/http"

type DecoratorFunc func(http.ResponseWriter, *http.Request, http.Handler) //패턴이 추상적인듯

type DecoHandler struct {
	fn DecoratorFunc
	h  http.Handler //인터페이스 embeded
}

func (self *DecoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { //자기꺼 구현
	self.fn(w, r, self.h)
}

func NewDecoHandler(h http.Handler, fn DecoratorFunc) http.Handler {
	return &DecoHandler{
		h:  h,  //mux
		fn: fn, //logger
	}
}
