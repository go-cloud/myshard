package topo

import (
	"fmt"
	"github.com/go-cloud/go-zookeeper/zk"
	"github.com/go-cloud/zkhelper"
	"path"
)

// agent:
func (t *Topo) GetAgentPath() string {
	return fmt.Sprintf("/zk/%s/agents", t.Name)
}

func (t *Topo) CreateAgent(nodeAddr string) (string, error) {
	zkhelper.CreateRecursive(t.conn, t.GetAgentPath(), "", 0, zkhelper.DefaultDirACLs())

	return t.conn.Create(path.Join(t.GetAgentPath(), nodeAddr), nil, zk.FlagEphemeral, zkhelper.DefaultFileACLs())
}

func (t *Topo) DeleteAgent(nodeAddr string) error {
	return t.conn.Delete(path.Join(t.GetAgentPath(), nodeAddr), -1)
}

func (t *Topo) ListAgents() ([]string, error) {
	as, _, err := t.conn.Children(t.GetAgentPath())
	return as, err
}
