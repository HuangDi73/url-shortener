package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func Chain(midds ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := range len(midds) - 1 {
			next = midds[i](next)
		}
		return next
	}
}
