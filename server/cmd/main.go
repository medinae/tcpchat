package main

import (
	"fmt"
	"log"
	"os"

	"github.com/medinae/tcpchat/server"
)

func main() {
	fmt.Println("Chat room will be started..")

	addr := getEnvOrDefault("TCPCHAT_SERVER_ADDR", ":8081")
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

func getEnvOrDefault(key, def string) string {
	v := os.Getenv(key)
	if "" == v {
		return def
	}
	return v
}
