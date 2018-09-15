package main

import (
	"flag"
	"fmt"

	"github.com/mitsukomegumi/Go-Rpcify/common"
	networkingTypes "github.com/mitsukomegumi/Go-Rpcify/networking/types"
	"github.com/mitsukomegumi/Go-Rpcify/types"
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

	server := networkingTypes.NewServer("server") // Initialize server

	call := helloWorldCall() // Initialize call

	server.Environment.AddCall(call) // Add hello world call

	err := server.StartServer() // Attempt to start server

	if err != nil {
		panic(err) // Panic
	}
}

func helloWorldCall() *types.Call {
	call, err := types.NewCall(common.HelloWorld, "HelloWorld") // Initialize call

	if err != nil { // Check for errors
		panic(err) // Panic
	}

	return call
}
