package handler

import (
	"net/http"

	gonic "github.com/gin-gonic/gin"
)

type healthResponse struct {
	Status string `json:"status"`
}

// Health responds with a minimal status payload.
func Health(c *gonic.Context) {
	c.JSON(http.StatusOK, healthResponse{Status: "ok"})
}
