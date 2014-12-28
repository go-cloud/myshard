package model

const (
	ProxyStatusUp   = "up"
	ProxyStatusDown = "down"
)

// Proxy is proxy server
type Proxy struct {
	// For MySQL protocol communication
	Addr string `json:"addr"`

	// Proxy HTTP communication
	HttpAddr string `json:"http_addr"`

	Status string `json:"status"`
}
