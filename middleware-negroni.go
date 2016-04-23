package kongojwt

import (
	"net/http"

	"github.com/gorilla/context"
)

type KongMiddleware struct {
}

type KongUser struct {
	KongID   string `json:"kong_id"`  // Kong internal ID
	Username string `json:"username"` // Username
	ID       string `json:"id"`       // Custom id
}

// Middleware is a struct that has a ServeHTTP method
func AuthMiddleware() *KongMiddleware {
	return &KongMiddleware{}
}

// The middleware handler
func (l *KongMiddleware) ServeHTTP(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	// Set information about current auth user into system
	user := KongUser{ID: req.Header.Get("X-Consumer-Custom-ID"), KongID: req.Header.Get("X-Consumer-ID"), Username: req.Header.Get("X-Consumer-Username")}
	context.Set(req, "auth", user)
	// Call the next middleware handler
	next(w, req)
}
