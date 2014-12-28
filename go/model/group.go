package model

const (
	NodeUpStatus   = "up"
	NodeDownStatus = "down"
)

const (
	NodeMasterType = "master"
	NodeSlaveType  = "slave"
)

// Node is a MySQL instance
type Node struct {
	Addr    string `json:"addr"`
	Type    string `json:"type"`
	Status  string `json:"status"`
	GroupID int    `json:"group_id"`
}

const (
	GroupUpStatus   = "up"
	GroupDownStatus = "down"
)

// Group includes one or more MySQL instances
// A group must have a master, none or more slaves
type Group struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
	Nodes  []Node `json:"nodes"`
}
