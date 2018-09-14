package types

import "github.com/mitsukomegumi/Go-Rpcify/types"

// Server - networking type used for rpc server
type Server struct {
	Endpoint    string             `json:"endpoint"`    // Used for http access to server
	Environment *types.Environment `json:"environment"` // Used for call access
}

// NewServer - initialize new instance of Server struct
func NewServer(endpoint string) *Server {
	env := types.NewEnvironment() // Initialize environment

	server := Server{endpoint, env} // Initialize server

	return &server // Return initialized server
}
