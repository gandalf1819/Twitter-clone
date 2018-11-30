package handler

import (
	"../services/auth/authpb"
	"google.golang.org/grpc"
	"log"
	"os"
)

type Config struct {
	Clients struct {
		AuthDB authpb.AuthTokenServiceClient
	}
	Connections struct {
		AuthToken *grpc.ClientConn
	}
	Addresses struct {
		AuthTokenPort string
	}
}

type Service int

const (
	AuthToken Service = 0
	User      Service = 1
	Post      Service = 2
)

var con Config

func (c *Config) RegisterClients() {
	authPort := os.Getenv("AUTH_PORT")
	if authPort == "" {
		panic("No Auth Port set in runBackendServer.sh file")
	}

	c.SetPortOfServices(AuthToken, authPort)

}

func (c *Config) SetPortOfServices(serviceType Service, port string) {
	switch serviceType {
	case AuthToken:
		c.Addresses.AuthTokenPort = port
		return
	}
}

func (c *Config) GetAuthTokenClient() authpb.AuthTokenServiceClient {
	return c.Clients.AuthDB
}

func (c *Config) DialServers() {
	options := grpc.WithInsecure()
	var err error
	c.Connections.AuthToken, err = grpc.Dial("localhost:"+c.Addresses.AuthTokenPort, options)
	if err != nil {
		log.Fatalf("could not connect to Auth Service: %v", err)
	} else {
		c.Clients.AuthDB = authpb.NewAuthTokenServiceClient(c.Connections.AuthToken)
		log.Println("SERVER: Successfully created a connection to Auth Service at", c.Addresses.AuthTokenPort)
	}
}

func InitializeConnectors() {
	con.RegisterClients()
	con.DialServers()
}
