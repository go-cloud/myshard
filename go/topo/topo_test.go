package topo

import (
	"flag"
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
