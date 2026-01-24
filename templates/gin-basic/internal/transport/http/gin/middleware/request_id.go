package middleware

import (
	"crypto/rand"
	"encoding/hex"
	"strconv"
	"time"

	gonic "github.com/gin-gonic/gin"
)

const requestIDHeader = "X-Request-Id"

// RequestID ensures every request has a traceable identifier.
func RequestID() gonic.HandlerFunc {
	return func(c *gonic.Context) {
		requestID := c.GetHeader(requestIDHeader)
		if requestID == "" {
			requestID = newRequestID()
		}

		c.Set("request_id", requestID)
		c.Writer.Header().Set(requestIDHeader, requestID)
		c.Next()
	}
}

func newRequestID() string {
	var buf [16]byte
	if _, err := rand.Read(buf[:]); err != nil {
		return strconv.FormatInt(time.Now().UnixNano(), 10)
	}
	return hex.EncodeToString(buf[:])
}
