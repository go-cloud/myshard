package model

const (
	ProxyStatusUp   = "up"
	ProxyStatusDown = "down"
)

// Proxy is proxy server
type Proxy struct {
	// For MySQL protocol communication
	Addr string

	// Proxy HTTP communication
	HttpAddr string

	Status string
}
