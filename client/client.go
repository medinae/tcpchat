package client

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
)

// ChatClient represent the contract to implement a
// chat client.
type ChatClient interface {
	ListenAndInteract()
}

// TCPChatClient represent a chat client based on TCP.
type TCPChatClient struct {
	con net.Conn
}

// NewTCPChatClient create a TCP connection to the given address and
// create a TCPChatClient struct pointer.
func NewTCPChatClient(addr string) (*TCPChatClient, error) {
	con, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &TCPChatClient{con}, nil
}

// ListenAndInteract listen for tcp message and process current user message to send.
func (cc *TCPChatClient) ListenAndInteract(cancel context.CancelFunc, username string) error {
	go cc.listen(cancel, username)
	for {
		ir := bufio.NewReader(os.Stdin)
		fmt.Print(username + ": ")
		msg, err := ir.ReadString('\n')
		if err != nil {
			return err
		}

		content := fmt.Sprintf("%s %s\n", username, msg)
		fmt.Fprintf(cc.con, content)
	}
}

func (cc *TCPChatClient) listen(cancel context.CancelFunc, username string) {
	for {
		r, err := bufio.NewReader(cc.con).ReadString('\n')
		if err != nil {
			fmt.Println("### Connection to server was lost... ###")
			cancel()
			break
		}

		cc.eraseCurrentLine()
		fmt.Print(r)
		fmt.Print(username + ": ")
	}
}

func (cc *TCPChatClient) eraseCurrentLine() {
	fmt.Printf("%c[2K", 27)
	fmt.Printf("\r")
}
