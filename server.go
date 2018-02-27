package main

import (
	"flag"
	"net"
	"time"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("server")

func main() {
	bind := flag.String("b", "0.0.0.0:8080", "bind an interface to accept tcp connection")
	createServer(*bind)
}

func createServer(bind string) {
	var (
		lis  *net.TCPListener
		addr *net.TCPAddr
		err  error
	)
	if addr, err = net.ResolveTCPAddr("tcp4", bind); err != nil {
		log.Errorf("net.ResolveTCPAddr(\"tcp4\", \"%s\") error(%v)", bind, err)
		return
	}
	if lis, err = net.ListenTCP("tcp4", addr); err != nil {
		log.Errorf("net.ListenTCP(\"tcp4\", \"%s\") error(%v)", bind, err)
		return
	}
	log.Infof("start tcp listen: \"%s\"", bind)
	for {
		var conn net.TCPConn
		if conn, err := lis.AcceptTCP(); err != nil {
			log.Error("listener.Accept(\"%s\") error(%v)", lis.Addr().String(), err)
			return
		}
		if err = conn.SetKeepAlive(true); err != nil {
			log.Error("conn.SetKeepAlive() error(%v)", err)
			time.Sleep(time.Second)
			continue
		}
		if err = conn.SetReadBuffer(1024); err != nil {
			log.Error("conn.SetReadBuffer() error(%v)", err)
			time.Sleep(time.Second)
			continue
		}
		if err = conn.SetWriteBuffer(4096); err != nil {
			log.Error("conn.SetWriteBuffer() error(%v)", err)
			time.Sleep(time.Second)
			continue
		}

	}
}
