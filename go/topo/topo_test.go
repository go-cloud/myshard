package topo

import (
	"flag"
	"github.com/go-cloud/zkhelper"
	. "gopkg.in/check.v1"
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
		if conn, err = zkhelper.ConnectToZk(*testZK); err != nil {
			c.Fatal(err)
		}
	}

	s.t = NewTopo("myshard_test", conn)

	if err := s.t.clear(); err != nil {
		c.Fatal(err)
	}
}

func (s *topoTestSuite) TearDownSuite(c *C) {
	s.t.Close()
}

func (s *topoTestSuite) TestAgent(c *C) {

}
