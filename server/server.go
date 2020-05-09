package server

import (
	"bufio"
	"fmt"
	"net"
)

// ChatServer represent the contract to implement a
// chat server.
type ChatServer interface {
	Close()
	Start() error
}

// TCPChatServer represent a chat server based on TCP.
type TCPChatServer struct {
	ln      net.Listener
	clients []net.Conn
}

func NewTCPChatServer(addr string) (*TCPChatServer, error) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &TCPChatServer{ln: ln}, err
}

func (ts *TCPChatServer) Close() {
	ts.ln.Close()
}

func (ts *TCPChatServer) Start() error {
	for {
		con, err := ts.ln.Accept()
		if err != nil {
			return err
		}
		ts.clients = append(ts.clients, con)
		fmt.Println("A new user just got connected..")
		go ts.handleCon(con)
	}
}

func (ts *TCPChatServer) handleCon(con net.Conn) {
	defer con.Close()
	for {
		r := bufio.NewReader(con)
		user, _ := r.ReadString(' ')
		msg, err := r.ReadString('\n')
		if err != nil {
			con.Write([]byte("nack\n"))
			continue
		}

		content := fmt.Sprintf("%s: %s", user, string(msg))
		ts.broadcastMessage(content, con)
	}
}

func (ts *TCPChatServer) broadcastMessage(content string, curCon net.Conn) {
	for _, con := range ts.clients {
		if curCon == con {
			// do not broadcast the message to the user that sent the message.
			continue
		}
		con.Write([]byte(content + "\n"))
	}
}
