package cors

import (
	"net/http"
	"strconv"
	"strings"
)

type Config struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	ExposeHeaders    []string
	AllowCredentials bool
	MaxAgeSeconds    int
}

func DefaultConfig() Config {
	return Config{}
}

func Middleware(cfg Config) func(http.Handler) http.Handler {
	allowMethods := strings.Join(cfg.AllowMethods, ", ")
	allowHeaders := strings.Join(cfg.AllowHeaders, ", ")
	exposeHeaders := strings.Join(cfg.ExposeHeaders, ", ")

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			if origin != "" && isOriginAllowed(origin, cfg.AllowOrigins) {
				applyHeaders(w, origin, cfg, allowMethods, allowHeaders, exposeHeaders)
				if isPreflight(r) {
					w.WriteHeader(http.StatusNoContent)
					return
				}
			} else if isPreflight(r) {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func isOriginAllowed(origin string, allowed []string) bool {
	if len(allowed) == 0 {
		return false
	}
	for _, entry := range allowed {
		if entry == "*" || strings.EqualFold(entry, origin) {
			return true
		}
	}
	return false
}

func isPreflight(r *http.Request) bool {
	return r.Method == http.MethodOptions && r.Header.Get("Access-Control-Request-Method") != ""
}

func applyHeaders(w http.ResponseWriter, origin string, cfg Config, allowMethods, allowHeaders, exposeHeaders string) {
	if contains(cfg.AllowOrigins, "*") && !cfg.AllowCredentials {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	} else {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Add("Vary", "Origin")
	}
	if allowMethods != "" {
		w.Header().Set("Access-Control-Allow-Methods", allowMethods)
	}
	if allowHeaders != "" {
		w.Header().Set("Access-Control-Allow-Headers", allowHeaders)
	}
	if exposeHeaders != "" {
		w.Header().Set("Access-Control-Expose-Headers", exposeHeaders)
	}
	if cfg.AllowCredentials {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	}
	if cfg.MaxAgeSeconds > 0 {
		w.Header().Set("Access-Control-Max-Age", intToString(cfg.MaxAgeSeconds))
	}
}

func contains(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

func intToString(value int) string {
	return strconv.Itoa(value)
}
