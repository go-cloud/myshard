package topo

import (
	"encoding/json"
	"fmt"
	"github.com/go-cloud/go-zookeeper/zk"
	"github.com/go-cloud/myshard/go/model"
	"github.com/go-cloud/zkhelper"
	"path"
)

func (t *Topo) GetProxyPath() string {
	return fmt.Sprintf("/zk/%s/proxys", t.Name)
}

func (t *Topo) CreateProxy(p *model.Proxy) (string, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return "", err
	}

	base := t.GetProxyPath()

	zkhelper.CreateRecursive(t.conn, base, "", 0, zkhelper.DefaultDirACLs())

	return t.conn.Create(path.Join(base, p.Addr), data, zk.FlagEphemeral, zkhelper.DefaultDirACLs())
}

func (t *Topo) DeleteProxy(addr string) error {
	return t.conn.Delete(path.Join(t.GetProxyPath(), addr), -1)
}

func (t *Topo) ListProxys() ([]*model.Proxy, error) {
	names, _, err := t.conn.Children(t.GetProxyPath())

	if err != nil && !zkhelper.ZkErrorEqual(err, zk.ErrNoNode) {
		return nil, err
	}

	ps := make([]*model.Proxy, 0, len(names))
	for _, n := range names {
		p, err := t.GetProxy(n)
		if err != nil {
			return nil, err
		}

		ps = append(ps, p)
	}

	return ps, nil
}

func (t *Topo) GetProxy(addr string) (*model.Proxy, error) {
	var p model.Proxy

	data, _, err := t.conn.Get(path.Join(t.GetProxyPath(), addr))
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &p); err != nil {
		return nil, err
	}

	return &p, nil
}

func (t *Topo) SetProxyStatus(addr string, status string) error {
	p, err := t.GetProxy(addr)
	if err != nil {
		return err
	}

	p.Status = status

	b, _ := json.Marshal(p)

	_, err = t.conn.Set(path.Join(t.GetProxyPath(), addr), b, -1)

	return err
}
