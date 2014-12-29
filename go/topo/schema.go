package topo

import (
	"encoding/json"
	"fmt"
	"github.com/go-cloud/myshard/go/model"
	"github.com/go-cloud/zkhelper"
	"path"
	"sort"
)

func (t *Topo) GetSchemaBasePath() string {
	return fmt.Sprintf("/zk/%s/schemas", t.Name)
}

func (t *Topo) GetSchemaPath(db string) string {
	return path.Join(t.GetSchemaBasePath(), db)
}

func (t *Topo) GetSchemaTablePath(db string, table string) string {
	return path.Join(t.GetSchemaPath(db), table)
}

func checkGroupIDExist(checkGroupID int, groupIDs []int) bool {
	found := false
	for _, groupID := range groupIDs {
		if groupID == checkGroupID {
			found = true
			break
		}
	}

	return found
}

func (t *Topo) CreateSchema(db string, groupIDs []int, defaultGroupID int) (string, error) {
	s := model.Schema{Name: db, GroupIDs: groupIDs, DefaultGroupID: defaultGroupID, Status: ""}

	sort.Ints(groupIDs)

	if !checkGroupIDExist(defaultGroupID, groupIDs) {
		return "", fmt.Errorf("group %d not in schema groups", defaultGroupID)
	}

	zkhelper.CreateRecursive(t.conn, t.GetSchemaBasePath(), "", 0, zkhelper.DefaultDirACLs())

	data, _ := json.Marshal(s)

	return t.conn.Create(t.GetSchemaPath(db), data, 0, zkhelper.DefaultDirACLs())
}

func (t *Topo) DeleteSchema(db string) error {
	return t.conn.Delete(t.GetSchemaPath(db), -1)
}

func (t *Topo) GetSchema(db string) (*model.Schema, error) {
	p := t.GetSchemaPath(db)
	data, _, err := t.conn.Get(p)
	if err != nil {
		return nil, err
	}

	var s model.Schema

	if err = json.Unmarshal(data, &s); err != nil {
		return nil, err
	}

	return &s, nil
}

func (t *Topo) UpdateSchemaDefaultGroup(db string, defaultGroupID int) error {
	s, err := t.GetSchema(db)
	if err != nil {
		return err
	}

	if !checkGroupIDExist(defaultGroupID, s.GroupIDs) {
		return fmt.Errorf("group %d not in schema groups", defaultGroupID)
	}

	s.DefaultGroupID = defaultGroupID

	data, _ := json.Marshal(s)

	_, err = t.conn.Set(t.GetSchemaPath(db), data, -1)
	return err
}

func (t *Topo) AddSchemaGroup(db string, groupID int) error {
	s, err := t.GetSchema(db)
	if err != nil {
		return err
	}

	if checkGroupIDExist(groupID, s.GroupIDs) {
		return nil
	}

	s.GroupIDs = append(s.GroupIDs, groupID)
	sort.Ints(s.GroupIDs)

	data, _ := json.Marshal(s)

	_, err = t.conn.Set(t.GetSchemaPath(db), data, -1)
	return err
}

func (t *Topo) DeleteSchemaGroup(db string, groupID int) error {
	s, err := t.GetSchema(db)
	if err != nil {
		return err
	}

	if !checkGroupIDExist(groupID, s.GroupIDs) {
		return fmt.Errorf("group %d not in schema groups", groupID)
	}

	groupIDs := make([]int, 0, len(s.GroupIDs))
	for _, id := range s.GroupIDs {
		if id != groupID {
			groupIDs = append(groupIDs, id)
		}
	}

	s.GroupIDs = groupIDs

	data, _ := json.Marshal(s)
	_, err = t.conn.Set(t.GetSchemaPath(db), data, -1)
	return err
}
