package main

import (
	"context"
	"fmt"
	"log"

	"github.com/medinae/tcpchat/client"
	"github.com/medinae/tcpchat/env"
)

func main() {
	addr := env.GetEnvOrDefault("TCPCHAT_SERVER_ADDR", ":8081")
	client, err := client.NewTCPChatClient(addr)
	if err != nil {
		log.Fatalf("error creating tcp chat client: %w", err)
	}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	var username string
	fmt.Print("Username:")
	fmt.Scanln(&username)

	client.ListenAndInteract(cancel, username)
}
