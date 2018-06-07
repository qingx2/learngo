package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRpc(host string, service interface{}) error {
	rpc.Register(service)

	listener, e := net.Listen("tcp", host)
	if e != nil {
		return e
	}
	log.Printf("Listening on %s", host)

	for {
		conn, e := listener.Accept()
		if e != nil {
			log.Printf("accept error: %v", e)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}

	return nil
}

func NewClient(host string) (*rpc.Client, error) {
	conn, e := net.Dial("tcp", host)
	if e != nil {
		return nil, e
	}

	return jsonrpc.NewClient(conn), nil
}
