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
	Addr    string
	Type    string
	Status  string
	GroupID int
}

const (
	GroupUpStatus   = "up"
	GroupDownStatus = "down"
)

// Group includes one or more MySQL instances
// A group must have a master, none or more slaves
type Group struct {
	ID     int
	Status string

	Nodes []Node
}
