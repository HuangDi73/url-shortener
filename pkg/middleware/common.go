package middleware

import "net/http"

type WriterWrapper struct {
	http.ResponseWriter
	StatusCode int
}

func (w *WriterWrapper) WriteHeader(StatusCode int) {
	w.ResponseWriter.WriteHeader(StatusCode)
	w.StatusCode = StatusCode
}
