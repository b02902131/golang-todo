package config

import (
	"net/http"

	"github.com/gorilla/handlers"
)

// SetupCORS 返回配置好的CORS中间件
func SetupCORS() func(http.Handler) http.Handler {
	return handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With"}),
		handlers.AllowedOrigins([]string{"*"}),
		// handlers.AllowedOrigins([]string{"http://localhost:8081"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
	)
}
