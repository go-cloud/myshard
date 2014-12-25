package main

import (
	"flag"
	log "github.com/golang/glog"
	"os"
	"os/signal"
	"syscall"
)

/*
   Agent monitors a MySQL server.

   Agent first create an ephemeral node in zookeeper(zk), and then pings MySQL every second. If the MySQL crashed,
   agent will close the connection with zk to let other services do something, like failover.

   If agent crashed but MySQL not, other services will think MySQL down too, but can not do failover. So we must also
   have a way to monitor agent, like using supervisor or crontab, the agent must be restarted quickly after down.

   You should put agent and MySQL in the same machine, if not, some network error may cause agent think MySQL was crashed.

   A MySQL server can only be monitored by an agent.
*/

var zkAddr = flag.String("zk", "127.0.0.1:2181", `zookeeper addres, you can use multi zk addresses seperated by ",", like 127.0.0.1:2181,127.0.0.1:2182`)
var myAddr = flag.String("mysql_addr", "127.0.0.1:3306", "local MySQL address")
var myUser = flag.String("mysql_user", "root", "local MySQL user")
var myPassword = flag.String("mysql_pass", "", "local MySQL password")

func main() {
	flag.Parse()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	<-c
	log.Info("ctrl-c or SIGTERM found, exit")
}
