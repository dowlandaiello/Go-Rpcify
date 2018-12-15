package main

import (
	"flag"
	"fmt"

	"github.com/dowlandaiello/Go-Rpcify/common"
	networkingTypes "github.com/dowlandaiello/Go-Rpcify/networking/types"
	"github.com/dowlandaiello/Go-Rpcify/types"
	"github.com/dowlandaiello/Go-Rpcify/upnp"
)

var (
	upnpFlag = flag.Bool("noupnp", true, "launch Go-Rpcify without upnp")
)

func main() {
	if *upnpFlag != true { // Check for noupnp flag
		fmt.Println("configuring upnp...") // Log upnp
		err := upnp.ForwardPort(8080)      // Attempt to forward server port

		if err != nil { // Check for errors
			fmt.Println(err.Error()) // Log error
		} else {
			upnp.ForwardPort(8081) // Attempt to forward RPC server port
		}
	}

	server := networkingTypes.NewServer("server") // Initialize server

	call := helloWorldCall() // Initialize call

	server.Environment.AddCall(call) // Add hello world call

	fmt.Println("-- SERVER -- starting server...")

	err := server.StartServer() // Attempt to start server

	if err != nil { // Check for errors
		panic(err) // Panic
	}
}

func helloWorldCall() *types.Call {
	console := new(common.Console) // Init console

	call, err := types.NewCall(console.HelloHTTP, console, "HelloWorld") // Initialize call

	fmt.Printf("initialized call with endpoint %s\n", call.Endpoint) // Log call

	if err != nil { // Check for errors
		panic(err) // Panic
	}

	return call // Return call
}
