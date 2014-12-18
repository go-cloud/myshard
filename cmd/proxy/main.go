package main

import (
	"flag"
	"fmt"
	"github.com/go-cloud/myshard/go/proxy"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var addr = flag.String("addr", "127.0.0.1:4000", "Proxy listen address")
var user = flag.String("user", "root", "Username")
var password = flag.String("password", "", "Password")
var usePprof = flag.Bool("pprof", false, "Enable pprof")
var pprofPort = flag.Int("pprof_port", 6060, "Pprof http port")

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()

	s, err := proxy.NewServer(*addr, *user, *password)
	if err != nil {
		println(err.Error())
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		os.Kill,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	if *usePprof {
		go func() {
			err := http.ListenAndServe(fmt.Sprintf(":%d", *pprofPort), nil)
			if err != nil {
				println(err.Error())
			}
		}()
	}

	<-sc

	s.Close()

}
