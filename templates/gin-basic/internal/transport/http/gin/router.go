package gin

import (
	gonic "github.com/gin-gonic/gin"

	"example.com/aitigo-gin-basic/internal/transport/http/gin/handler"
)

// Register wires HTTP routes for the API.
func Register(router *gonic.Engine) {
	router.GET("/health", handler.Health)
}
