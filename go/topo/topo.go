package topo

import (
	"fmt"
	"github.com/go-cloud/go-zookeeper/zk"
	"github.com/go-cloud/zkhelper"
)

// We use zookeeper to store the configuration and cooperate
// all services
type Topo struct {
	// for the base zk directory name
	Name string

	conn zkhelper.Conn
}

func NewTopo(name string, conn zkhelper.Conn) *Topo {
	t := new(Topo)
	t.Name = name
	t.conn = conn
	return t
}

func (t *Topo) Close() {
	t.conn.Close()
}

// Clear all zk data, very dangerous to use
func (t *Topo) clear() error {
	p := fmt.Sprintf("/zk/%s", t.Name)

	err := zkhelper.DeleteRecursive(t.conn, p, -1)
	if zkhelper.ZkErrorEqual(err, zk.ErrNoNode) {
		return nil
	} else {
		return err
	}
}
