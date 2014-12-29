package topo

import (
	"encoding/json"
	"fmt"
	"github.com/go-cloud/myshard/go/model"
	"github.com/go-cloud/zkhelper"
	"path"
	"sort"
)

func (t *Topo) GetGroupBasePath() string {
	return fmt.Sprintf("/zk/%s/groups", t.Name)
}

func (t *Topo) GetGroupPath(groupID int) string {
	return path.Join(t.GetGroupBasePath(), fmt.Sprintf("group_%06d", groupID))
}

func (t *Topo) GetGroupNodePath(groupID int, addr string) string {
	return path.Join(t.GetGroupPath(groupID), addr)
}

// group:
//  /zk/myshard/groups/group_000001
//  /zk/myshard/groups/group_000002
func (t *Topo) CreateGroup(groupID int) (string, error) {
	group := model.Group{groupID, model.GroupUpStatus, nil}

	zkhelper.CreateRecursive(t.conn, t.GetGroupBasePath(), "", 0, zkhelper.DefaultDirACLs())

	data, _ := json.Marshal(group)

	return t.conn.Create(t.GetGroupPath(groupID), data, 0, zkhelper.DefaultDirACLs())
}

func (t *Topo) DeleteGroup(groupID int) error {
	return t.conn.Delete(t.GetGroupPath(groupID), -1)
}

func (t *Topo) SetGroupStatus(groupID int, status string) error {
	group, err := t.getGroup(groupID)
	if err != nil {
		return err
	}

	group.Status = status

	data, _ := json.Marshal(group)

	_, err = t.conn.Set(t.GetGroupPath(groupID), data, -1)
	return err
}

func (t *Topo) GetGroup(groupID int) (*model.Group, error) {
	group, err := t.getGroup(groupID)
	if err != nil {
		return nil, err
	}

	p := t.GetGroupPath(groupID)
	children, _, err := t.conn.Children(p)
	if err != nil {
		return nil, err
	}

	group.Nodes = make([]model.Node, 0, len(children))

	var data []byte
	for _, child := range children {
		if data, _, err = t.conn.Get(path.Join(p, child)); err != nil {
			return nil, err
		}

		var node model.Node

		if err = json.Unmarshal(data, &node); err != nil {
			return nil, err
		}

		group.Nodes = append(group.Nodes, node)
	}

	sort.Sort(model.NodeSlice(group.Nodes))

	return group, nil
}

func (t *Topo) getGroup(groupID int) (*model.Group, error) {
	var group model.Group

	p := t.GetGroupPath(groupID)

	data, _, err := t.conn.Get(p)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &group); err != nil {
		return nil, err
	}

	return &group, nil
}

func (t *Topo) AddGroupNode(groupID int, addr string, tp string, status string, weight int) (string, error) {
	node := model.Node{GroupID: groupID, Addr: addr, Type: tp, Status: status, Weight: weight}

	data, _ := json.Marshal(node)

	return t.conn.Create(t.GetGroupNodePath(groupID, addr), data, 0, zkhelper.DefaultDirACLs())
}

func (t *Topo) SetGroupNodeStatus(groupID int, addr string, status string) error {
	p := t.GetGroupNodePath(groupID, addr)
	node, err := t.GetGroupNode(groupID, addr)
	if err != nil {
		return err
	}

	node.Status = status
	data, _ := json.Marshal(node)

	_, err = t.conn.Set(p, data, -1)
	return err
}

func (t *Topo) SetGroupNodeWeight(groupID int, addr string, weight int) error {
	p := t.GetGroupNodePath(groupID, addr)
	node, err := t.GetGroupNode(groupID, addr)
	if err != nil {
		return err
	}

	node.Weight = weight
	data, _ := json.Marshal(node)

	_, err = t.conn.Set(p, data, -1)
	return err
}

func (t *Topo) DeleteGroupNode(groupID int, addr string) error {
	return t.conn.Delete(t.GetGroupNodePath(groupID, addr), -1)
}

func (t *Topo) GetGroupNode(groupID int, addr string) (*model.Node, error) {
	p := t.GetGroupNodePath(groupID, addr)
	data, _, err := t.conn.Get(p)
	if err != nil {
		return nil, err
	}

	var node model.Node
	if err = json.Unmarshal(data, &node); err != nil {
		return nil, err
	}

	return &node, nil
}
