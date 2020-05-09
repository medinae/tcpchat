package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
)

func main() {
	con, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	var username string
	fmt.Print("Username:")
	fmt.Scanln(&username)

	go listen(cancel, con, username)
	for {
		ir := bufio.NewReader(os.Stdin)
		fmt.Print(username + ": ")
		msg, err := ir.ReadString('\n')
		if err != nil {
			panic(err)
		}

		content := fmt.Sprintf("%s %s\n", username, msg)
		fmt.Fprintf(con, content)
	}
}

func listen(cancel context.CancelFunc, con net.Conn, username string) {
	for {
		r, err := bufio.NewReader(con).ReadString('\n')
		if err != nil {
			fmt.Println("### Connection to server was lost... ###")
			cancel()
			panic(err)
		}

		eraseCurrentLine()
		fmt.Print(r)
		fmt.Print(username + ": ")
	}
}

func eraseCurrentLine() {
	fmt.Printf("%c[2K", 27)
	fmt.Printf("\r")
}
