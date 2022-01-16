package internal

import (
	"github.com/mytord/fs/gateway/internal/middleware"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type ApiGateway struct {
	origin    *url.URL
	jwtSecret string
}

func NewApiGateway(origin *url.URL, jwtSecret string) *ApiGateway {
	return &ApiGateway{
		origin:    origin,
		jwtSecret: jwtSecret,
	}
}

func (gateway *ApiGateway) Handle(w http.ResponseWriter, r *http.Request) {
	proxy := httputil.NewSingleHostReverseProxy(gateway.origin)

	auth := middleware.NewAuth(gateway.jwtSecret, time.Hour*12)

	// request middlewares
	var director middleware.RequestMiddleware
	director = proxy.Director
	director = auth.Authorize(director)
	proxy.Director = director

	// response middlewares
	var modifyResponse middleware.ResponseMiddleware
	modifyResponse = auth.ApplyToken()
	proxy.ModifyResponse = modifyResponse

	proxy.ServeHTTP(w, r)
}
