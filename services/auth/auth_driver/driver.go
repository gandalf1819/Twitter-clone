package auth_driver

import (
	"../authpb"
	"context"
	"errors"
	"log"
	"math/rand"
)

type Server struct{}

type TokenDB struct {
	t *authpb.AuthToken
}

var db TokenDB

func Init() {
	tok := NewToken()
	db.t = tok
	log.Println("DB Token Initialized =", db.t.Token)
}

func NewToken() *authpb.AuthToken {
	return &authpb.AuthToken{
		Token: make(map[string]int32),
	}
}

func (*Server) GetUserIdFromToken(ctx context.Context, tokName *authpb.AuthTokenName) (*authpb.AuthTokenValue, error) {
	token := &authpb.AuthTokenValue{}
	userId, ok := db.t.Token[tokName.TokenName]
	if ok {
		token.TokenValue = userId
		return token, nil
	} else {
		return token, errors.New("No token found or token expired!!")
	}
}

func (s *Server) AddToken(ctx context.Context, userId *authpb.UserId) (*authpb.AuthTokenName, error) {
	tok, err := s.GenerateToken(ctx, &authpb.InitToken{})

	if err != nil {
		log.Println("Error occured while adding token for userId", userId.User, " = ", err)
		return &authpb.AuthTokenName{TokenName: ""}, err
	}

	db.t.Token[tok.TokenName] = userId.User

	log.Println("Token generated for userId ", userId.User, " = ", tok.TokenName)
	log.Println("TokenDB = ", db.t.Token)

	return tok, err
}

func (*Server) UnsetToken(ctx context.Context, tokName *authpb.AuthTokenName) (*authpb.Status, error) {
	status := &authpb.Status{}
	delete(db.t.Token, tokName.TokenName)
	status.ResponseStatus = true
	log.Println("Token ", tokName.TokenName, " deleted successfully")
	log.Println("TokenDB = ", db.t.Token)
	return status, nil
}

func (*Server) GenerateToken(ctx context.Context, in *authpb.InitToken) (*authpb.AuthTokenName, error) {
	letterBytes := "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, 16)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	status := &authpb.AuthTokenName{
		TokenName: string(b),
	}

	log.Println("Token Generated = ", status.TokenName)

	return status, nil
}
