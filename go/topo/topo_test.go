package topo

import (
	"flag"
	"github.com/go-cloud/myshard/go/model"
	"github.com/go-cloud/zkhelper"
	. "gopkg.in/check.v1"
	"path"
	"sort"
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
	c.Assert(s1, Equals, path.Join(s.t.GetAgentPath(), agent1))

	s2, err := s.t.CreateAgent(agent2)
	c.Assert(err, IsNil)
	c.Assert(s2, Equals, path.Join(s.t.GetAgentPath(), agent2))

	as, err := s.t.ListAgents()
	c.Assert(err, IsNil)

	sort.Strings(as)

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
