package model

// Schema is MySQL database, and can be stored in one or more groups.
type Schema struct {
	Name string

	GroupIDs []int

	// The default group for the table stored
	DefaultGroupID int

	// Later use, todo, maybe for migrate
	Status string
}

type Table struct {
	Schema string

	// num > 0 If table is spllited
	SubTableNum int

	// If the table is not splitted, it will be stored in this group
	GroupID int

	// If the table is splitted, e.g. a user table, may be splitted into 1024,
	// so the real table may be stored in different groups
	// like user_0 stored in group_0, user_1 stored in group_1
	// key format:
	//  1        -> user_1
	//  1,2,3    -> user_1, user_2 and user_3
	//  1-3      -> user_1, user_2 and user_3
	// value is the group id
	SubTables map[string]int

	// Later use, todo, maybe for migrate
	Status string
}
