# Go-Rpcify

A Go library for generating event-based RPC endpoints.

## Installation

Installation via Go Get:

```BASH
go get github.com/mitsukomegumi/Go-Rpcify
```

## Usage

### Setup

Initialize an RPC Server:

```Go
import (
    serverTypes github.com/mitsukomegumi/Go-Rpcify/networking/types
)

server := serverTypes.NewServer("server") // Initialize server
```

Initialize a Call from a Method:

```Go
import (
    github.com/mitsukomegumi/Go-Rpcify/types
)

call, err := types.NewCall(method, "methodName") // Initialize call

if err != nil { // Check for errors
    panic(err) // Panic
}
```

Add a Call to an RPC Server:

```Go
server.Environment.AddCall(call) // Add call
```

Start a Server:

```Go
server.StartServer() // Start Server
```

### Method Support

In order to support the standard go net/rpc RPC server implementation, all call methods must

1. Be of an exported type

2. Be exported

3. Have two arguments (an arguments list struct, and a response pointer)

4. Have a return type of error

5. Have a mirroring standalone method with no arguments, and a return type of string and error (response, and error)

### Server Interaction

Connect to a Server via RPC:

```RPC
client, err := rpc.DialHTTP("tcp", "localhost" + ":8081") // Connect to server

if err != nil { // Check for errors
    panic(err) // Panic
}
```

Run a Call via RPC:

```RPC
var reply int // Init reply buffer

err = client.Call("Common.HelloWorld", &reply) // Call method

if err != nil { // Check for errors
    panic(err) // Panic
}
```

Run a Call via HTTP:

```HTTP
http://localhost:8080/call/$callEndpoint
```

Run Go via HTTP:

```HTTP
http://localhost:8080/call/$goCode
```

    Example:

    http://localhost:8080/call/fmt.Println("example")
