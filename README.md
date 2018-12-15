# Go-Rpcify

A Go library for generating event-based RPC endpoints.

## Installation

Installation via Go Get:

```BASH
go get github.com/dowlandaiello/Go-Rpcify
```

## Usage

### Setup

Initialize an RPC Server:

```Go
import (
    serverTypes github.com/dowlandaiello/Go-Rpcify/networking/types
)

server := serverTypes.NewServer("server") // Initialize server
```

Initialize a Call from a Method:

```Go
import (
    github.com/dowlandaiello/Go-Rpcify/types
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

### Server Interaction

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
