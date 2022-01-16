package middleware

import "net/http"

type RequestMiddleware func(*http.Request)
type ResponseMiddleware func(*http.Response) error
