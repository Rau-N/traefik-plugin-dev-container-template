package traefik_plugin_hello

import (
	"context"
	"net/http"
)

type Config struct {
	HeaderKey   string `json:"headerKey,omitempty" toml:"headerKey,omitempty" yaml:"headerKey,omitempty"`
	HeaderValue string `json:"headerValue,omitempty" toml:"headerValue,omitempty" yaml:"headerValue,omitempty"`
}

func CreateConfig() *Config {
	return &Config{
		HeaderKey:   "X-Plugin",
		HeaderValue: "Hello",
	}
}

type hello struct {
	next  http.Handler
	name  string
	key   string
	value string
}

func New(_ context.Context, next http.Handler, cfg *Config, name string) (http.Handler, error) {
	return &hello{
		next:  next,
		name:  name,
		key:   cfg.HeaderKey,
		value: cfg.HeaderValue,
	}, nil
}

func (h *hello) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set(h.key, h.value)
	h.next.ServeHTTP(rw, req)
}
