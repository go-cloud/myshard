package model

// Schema is MySQL database, and can be stored in one or more groups.
type Schema struct {
	Name string `json:"name"`

	GroupIDs []int `json:"group_ids"`

	// The default group for the table stored
	DefaultGroupID int `json:"default_group_id"`

	// Later use, todo, maybe for migrate
	Status string `json:"status"`
}

type Table struct {
	DB string `json:"db"`

	// If the table is not splitted, it will be stored in this group
	GroupID int `json:"group_id"`

	// If the table is splitted, e.g. a user table, may be splitted into 1024,
	// so the real table may be stored in different groups
	// like user_0 stored in group_0, user_1 stored in group_1
	// the sub table id must start from 0, so we can use a slice to map
	// the table id -> group id
	SubTables []int `json:"sub_tables"`

	// Later use, todo, maybe for migrate
	Status string `json:"status"`
}
