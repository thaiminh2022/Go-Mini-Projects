package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func CreateStack(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			var x Middleware = xs[i]
			next = x(next)

		}
		return next
	}
}
