# Go-Rpcify

A Go library for generating event-based RPC endpoints.

## Installation

Installation via Go Get:

```BASH
go get github.com/mitsukomegumi/Go-Rpcify
```

## Usage

Initialize an RPC Server:

```Go
server := networkingTypes.NewServer("server") // Initialize server
```

Initialize a Call from a Method:

```Go
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