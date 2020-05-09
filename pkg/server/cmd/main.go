package main

import (
	"fmt"
	"log"

	"github.com/medinae/tcpchat/pkg/env"
	"github.com/medinae/tcpchat/pkg/server"
)

func main() {
	fmt.Println("Chat room will be started..")

	addr := env.GetEnvOrDefault("TCPCHAT_SERVER_ADDR", ":8081")
	chat, err := server.NewTCPChatServer(addr)
	if err != nil {
		log.Fatalf("error creating tcp chat server: %w", err)
	}
	defer chat.Close()

	err = chat.Start()
	if err != nil {
		log.Fatalf("error starting the chat server: %w", err)
	}
}
