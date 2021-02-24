package util

import (
	"net/http"
	"strings"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
)

type GrpcWebMiddleware struct {
	*grpcweb.WrappedGrpcServer
}

func (m *GrpcWebMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Upgrade", strings.ToLower(r.Header.Get("Upgrade")))
		if m.IsAcceptableGrpcCorsRequest(r) || m.IsGrpcWebRequest(r) || m.IsGrpcWebSocketRequest(r) {
			r.Header.Set("Wsauthtoken", r.URL.Query().Get("token"))
			m.ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func NewGrpcWebMiddleware(grpcWeb *grpcweb.WrappedGrpcServer) *GrpcWebMiddleware {
	return &GrpcWebMiddleware{grpcWeb}
}

func (m *GrpcWebMiddleware) DefaultFailureHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
}
