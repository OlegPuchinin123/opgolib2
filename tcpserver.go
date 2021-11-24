package opgolib2

import (
	"fmt"
	"net"
)

type TCPServer struct {
	M_ln    *net.TCPListener
	M_laddr *net.TCPAddr
}

type ServerClient struct {
	M_conn *net.TCPConn
}

func (serv *TCPServer) NewTCPServer() *TCPServer {
	return &TCPServer{}
}

// n - network type
func (serv *TCPServer) Listen(n string, host string, port uint16) error {
	var (
		e        error
		hostport string
	)

	hostport = fmt.Sprintf("[%s]:%d", host, port)
	serv.M_laddr, e = net.ResolveTCPAddr(n, hostport)
	serv.M_ln, e = net.ListenTCP(n, serv.M_laddr)
	return e
}

func (serv *TCPServer) Accept() (*ServerClient, error) {
	var (
		e  error
		sc *ServerClient
	)
	sc = &ServerClient{}
	sc.M_conn, e = serv.M_ln.AcceptTCP()
	if e != nil {
		return nil, e
	}
	return sc, nil
}

func (sc *ServerClient) Recv(buf []byte) (int, error) {
	return sc.M_conn.Read(buf)
}

func (sc *ServerClient) Send(buf []byte) (int, error) {
	return sc.M_conn.Write(buf)
}

func (sc *ServerClient) Close() {
	sc.M_conn.Close()
}
