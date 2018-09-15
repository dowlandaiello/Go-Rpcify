package main

import (
	"flag"
	"fmt"

	"github.com/mitsukomegumi/Go-Rpcify/networking/types"
	"github.com/mitsukomegumi/Go-Rpcify/upnp"
)

var (
	upnpFlag = flag.Bool("noupnp", false, "launch Go-Rpcify without upnp")
)

func main() {
	if !*upnpFlag {
		err := upnp.ForwardPort(8080) // Attempt to forward server port

		if err != nil { // Check for errors
			fmt.Println(err.Error()) // Log error
		}
	}

	server := types.NewServer("server") // Initialize server

	err := server.StartServer() // Attempt to start server

	if err != nil {
		panic(err) // Panic
	}
}
