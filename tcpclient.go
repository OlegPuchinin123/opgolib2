package opgolib2

import (
	"fmt"
	"net"
)

type TCPClient struct {
	M_conn  *net.TCPConn
	M_raddr *net.TCPAddr
}

func NewTCPClient() *TCPClient {
	return &TCPClient{}
}

// n - network type
func (c *TCPClient) Connect(n string, host string, port uint16) error {
	var (
		e        error
		hostport string
	)

	hostport = fmt.Sprintf("[%s]:%d", host, port)
	c.M_raddr, e = net.ResolveTCPAddr(n, hostport)
	if e != nil {
		return e
	}
	//log.Default().Printf("Resolve OK")
	c.M_conn, e = net.DialTCP(n, nil, c.M_raddr)
	if e != nil {
		return e
	}
	//log.Default().Printf("Dial OK")
	return nil
}

func (c *TCPClient) Send(buf []byte) (int, error) {
	var (
		e error
		n int
	)
	n, e = c.M_conn.Write(buf)
	return n, e
}

func (c *TCPClient) Recv(buf []byte) (int, error) {
	var (
		e error
		n int
	)
	n, e = c.M_conn.Read(buf)
	return n, e
}

func (c *TCPClient) Close() {
	c.M_conn.Close()
}
