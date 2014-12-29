package topo

import (
	"flag"
	"github.com/go-cloud/myshard/go/model"
	"github.com/go-cloud/zkhelper"
	. "gopkg.in/check.v1"
	"path"
	"testing"
)

var testZK = flag.String("zk", "", "Use zookeeper for test if set")

func Test(t *testing.T) {
	TestingT(t)
}

type topoTestSuite struct {
	t *Topo
}

var _ = Suite(&topoTestSuite{})

func (s *topoTestSuite) SetUpSuite(c *C) {
	var conn zkhelper.Conn

	if len(*testZK) == 0 {
		//use fake zk
		conn = zkhelper.NewConn()
	} else {
		var err error
		conn, err = zkhelper.ConnectToZk(*testZK)
		c.Assert(err, IsNil)
	}

	s.t = NewTopo("myshard_test", conn)

	err := s.t.clear()
	c.Assert(err, IsNil)
}

func (s *topoTestSuite) TearDownSuite(c *C) {
	s.t.Close()
}

func (s *topoTestSuite) TestAgent(c *C) {
	agent1 := "127.0.0.1:3306"
	agent2 := "127.0.0.1:3307"

	s1, err := s.t.CreateAgent(agent1)
	c.Assert(err, IsNil)
	c.Assert(s1, Equals, path.Join(s.t.GetAgentBasePath(), agent1))

	s2, err := s.t.CreateAgent(agent2)
	c.Assert(err, IsNil)
	c.Assert(s2, Equals, path.Join(s.t.GetAgentBasePath(), agent2))

	as, err := s.t.ListAgents()
	c.Assert(err, IsNil)

	c.Assert(as, DeepEquals, []string{agent1, agent2})
}

func (s *topoTestSuite) TestProxy(c *C) {
	p1 := &model.Proxy{"127.0.0.1:4000", "127.0.0.1:14000", "up"}
	p2 := &model.Proxy{"127.0.0.1:4001", "127.0.0.1:14001", "up"}

	_, err := s.t.CreateProxy(p1)
	c.Assert(err, IsNil)

	_, err = s.t.CreateProxy(p2)
	c.Assert(err, IsNil)

	ps, err := s.t.ListProxys()
	c.Assert(err, IsNil)
	c.Assert(ps, HasLen, 2)

	err = s.t.SetProxyStatus(p1.Addr, "down")
	c.Assert(err, IsNil)

	p, err := s.t.GetProxy(p1.Addr)
	c.Assert(err, IsNil)

	c.Assert(p.Status, Equals, "down")
}

func (s *topoTestSuite) TestGroup(c *C) {
	_, err := s.t.CreateGroup(1)
	c.Assert(err, IsNil)

	_, err = s.t.AddGroupNode(1, "127.0.0.1:3306", model.NodeMasterType, model.NodeDownStatus, 0)
	c.Assert(err, IsNil)

	_, err = s.t.AddGroupNode(1, "127.0.0.1:3307", model.NodeSlaveType, model.NodeDownStatus, 0)
	c.Assert(err, IsNil)

	err = s.t.SetGroupNodeStatus(1, "127.0.0.1:3306", model.NodeUpStatus)
	c.Assert(err, IsNil)

	err = s.t.SetGroupNodeStatus(1, "127.0.0.1:3307", model.NodeUpStatus)
	c.Assert(err, IsNil)

	g, err := s.t.GetGroup(1)
	c.Assert(err, IsNil)
	c.Assert(g.ID, Equals, 1)
	c.Assert(g.Nodes, DeepEquals, []model.Node{
		model.Node{1, "127.0.0.1:3306", model.NodeMasterType, model.NodeUpStatus, 0},
		model.Node{1, "127.0.0.1:3307", model.NodeSlaveType, model.NodeUpStatus, 0},
	})

	err = s.t.DeleteGroup(1)
	c.Assert(err, NotNil)

	err = s.t.DeleteGroupNode(1, "127.0.0.1:3306")
	c.Assert(err, IsNil)

	err = s.t.DeleteGroupNode(1, "127.0.0.1:3307")
	c.Assert(err, IsNil)

	err = s.t.DeleteGroup(1)
	c.Assert(err, IsNil)
}
