package model

const (
	RouteTypeHash  = "hash"
	RouteTypeRange = "range"
)

// We will route to let a SQL exexuted in the exact group
type Route struct {
	Schema string
	Table  string
	Column string
	Type   string

	// For range type
	// key format 1-10000, left close and right open
	// value is sub table id
	// e.g, if the ranges is {"0-100" : 0, "100-200" : 1}
	// if a key is 50, we may know that the data will be saved in table_0
	Ranges map[string]int
}
