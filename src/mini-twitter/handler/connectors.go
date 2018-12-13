package handler

import (
	"google.golang.org/grpc"
	"log"
	"mini-twitter/services/auth/authpb"
	"mini-twitter/services/post/postpb"
	"mini-twitter/services/user/userpb"
	"os"
)

type Config struct {
	Clients struct {
		AuthDB     authpb.AuthTokenServiceClient
		UserPostDB postpb.PostServiceClient
		UserDB     userpb.UserServiceClient
	}
	Connections struct {
		AuthToken *grpc.ClientConn
		UserPost  *grpc.ClientConn
		User      *grpc.ClientConn
	}
	Addresses struct {
		AuthTokenPort string
		UserPostPort  string
		UserPort      string
	}
}

type Service int

const (
	AuthToken Service = 0
	User      Service = 1
	UserPost  Service = 2
)

var con Config

func (c *Config) RegisterClients() {
	authPort := os.Getenv("AUTH_PORT")
	postPort := os.Getenv("USER_POST_PORT")
	userPort := os.Getenv("USER_PORT")
	if authPort == "" {
		panic("No Auth Port set in runBackendServer.sh file")
	}
	if postPort == "" {
		panic("No Post Port set in runBackendServer.sh file")
	}
	if userPort == "" {
		panic("No User Port set in runBackendServer.sh file")
	}

	c.SetPortOfServices(AuthToken, authPort)
	c.SetPortOfServices(UserPost, postPort)
	c.SetPortOfServices(User, userPort)

}

func (c *Config) SetPortOfServices(serviceType Service, port string) {
	switch serviceType {
	case AuthToken:
		c.Addresses.AuthTokenPort = port
		return
	case UserPost:
		c.Addresses.UserPostPort = port
		return
	case User:
		c.Addresses.UserPort = port
		return
	}
}

func (c *Config) GetAuthTokenClient() authpb.AuthTokenServiceClient {
	return c.Clients.AuthDB
}

func (c *Config) GetUserPostClient() postpb.PostServiceClient {
	return c.Clients.UserPostDB
}

func (c *Config) GetUserClient() userpb.UserServiceClient {
	return c.Clients.UserDB
}

func (c *Config) DialServers() {
	options := grpc.WithInsecure()
	var err error
	//Token Client
	c.Connections.AuthToken, err = grpc.Dial("localhost:"+c.Addresses.AuthTokenPort, options)
	if err != nil {
		log.Fatalf("could not connect to Auth Service: %v", err)
	} else {
		c.Clients.AuthDB = authpb.NewAuthTokenServiceClient(c.Connections.AuthToken)
		log.Println("SERVER: Successfully created a connection to Auth Service at", c.Addresses.AuthTokenPort)
	}

	//UserPosts Client
	c.Connections.UserPost, err = grpc.Dial("localhost:"+c.Addresses.UserPostPort, options)
	if err != nil {
		log.Fatalf("could not connect to UserPost Service: %v", err)
	} else {
		c.Clients.UserPostDB = postpb.NewPostServiceClient(c.Connections.UserPost)
		log.Println("SERVER: Successfully created a connection to User Post Service at", c.Addresses.UserPostPort)
	}

	//User Client
	c.Connections.User, err = grpc.Dial("localhost:"+c.Addresses.UserPort, options)
	if err != nil {
		log.Fatalf("could not connect to User Service: %v", err)
	} else {
		c.Clients.UserDB = userpb.NewUserServiceClient(c.Connections.User)
		log.Println("SERVER: Successfully created a connection to User Service at", c.Addresses.UserPort)
	}
}

func InitializeConnectors() {
	con.RegisterClients()
	con.DialServers()
}
