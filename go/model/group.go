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
	GroupID int    `json:"group_id"`
	Addr    string `json:"addr"`
	Type    string `json:"type"`
	Status  string `json:"status"`
	Weight  int    `json:"weight"`
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
	Nodes  []Node `json:"-"`
}

type NodeSlice []Node

func (ns NodeSlice) Len() int {
	return len(ns)
}

// master node is at the beginning
// for same type, higher weight is in front
func (ns NodeSlice) Less(i, j int) bool {
	if ns[i].Type != ns[j].Type {
		if ns[i].Type == NodeMasterType {
			return true
		} else {
			return false
		}
	} else {
		return ns[i].Weight >= ns[j].Weight
	}
}

func (ns NodeSlice) Swap(i, j int) {
	ns[i] = ns[j]
}
