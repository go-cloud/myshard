package topo

import (
	"fmt"
	"github.com/go-cloud/go-zookeeper/zk"
	"github.com/go-cloud/zkhelper"
	"path"
	"sort"
)

// agent:
func (t *Topo) GetAgentBasePath() string {
	return fmt.Sprintf("/zk/%s/agents", t.Name)
}

func (t *Topo) CreateAgent(nodeAddr string) (string, error) {
	zkhelper.CreateRecursive(t.conn, t.GetAgentBasePath(), "", 0, zkhelper.DefaultDirACLs())

	return t.conn.Create(path.Join(t.GetAgentBasePath(), nodeAddr), nil, zk.FlagEphemeral, zkhelper.DefaultFileACLs())
}

func (t *Topo) DeleteAgent(nodeAddr string) error {
	return t.conn.Delete(path.Join(t.GetAgentBasePath(), nodeAddr), -1)
}

func (t *Topo) ListAgents() ([]string, error) {
	as, _, err := t.conn.Children(t.GetAgentBasePath())
	sort.Strings(as)
	return as, err
}
