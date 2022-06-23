package server

import (
	"net/http"
	"runtime/debug"

	pctx "github.com/appootb/substratum/v2/plugin/context"
)

type handlerWrapper struct {
	component string
	handler   http.Handler
}

func (h *handlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if result := recover(); result != nil {
			debug.PrintStack()
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()
	ctx := pctx.WithImplementContext(r.Context(), h.component)
	h.handler.ServeHTTP(w, r.WithContext(ctx))
}

type httpServeMux struct {
	component string
	serveMux  *http.ServeMux
}

func (h *httpServeMux) Handle(pattern string, handler http.Handler) {
	h.serveMux.Handle(pattern, &handlerWrapper{
		component: h.component,
		handler:   handler,
	})
}

func (h *httpServeMux) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	h.serveMux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if result := recover(); result != nil {
				debug.PrintStack()
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		ctx := pctx.WithImplementContext(r.Context(), h.component)
		handler(w, r.WithContext(ctx))
	})
}
