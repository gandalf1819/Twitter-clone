package auth_driver

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"io/ioutil"
	"log"
	"math/rand"
	"mini-twitter/services/auth/authpb"
	"net/http"
	"strings"
)

type Server struct{}

type TokenDB struct {
	t *authpb.AuthToken
}

var tokenDB authpb.AuthToken

func Init() {
	var db TokenDB
	tok := NewToken()
	db.t = tok

	_, err := InteractWithRaftStorage("PUT", "tokenDB", db.t)
	if err != nil {
		log.Println("Error occured while storing token data in Raft =", err)
		panic(err)
	}

	log.Println("DB Token Initialized =", db.t.Token)
}

func GetTokenDB(value interface{}) (*authpb.AuthToken, error) {
	var db TokenDB
	data, err := InteractWithRaftStorage("GET", "tokenDB", db.t)
	if err != nil {
		log.Println("Error occured while getting token data from Raft =", err)
		panic(err)
	}
	var tokenDB *authpb.AuthToken
	tokenDB, err = DecodeRaftTokenStorage(data)
	if err != nil {
		log.Println("Error occured while decoding token data from Raft storage =", err)
		return nil, err
	}
	log.Println("tokenDB after decode =", tokenDB)
	return tokenDB, nil
}

func DecodeRaftTokenStorage(db string) (*authpb.AuthToken, error) {
	log.Println("Decode Token Storage called")
	dec := gob.NewDecoder(bytes.NewBufferString(db))
	if err := dec.Decode(&tokenDB); err != nil {
		log.Fatalf("raftexample: could not decode message (%v)", err)
		return nil, err
	}
	log.Println("tokenDB in DecodeRaftTokenStorage =", tokenDB)

	return &tokenDB, nil
}

func InteractWithRaftStorage(method string, key string, value interface{}) (string, error) {
	log.Println("Interacted with Raft, method called =", method)
	var payloadValue string
	if method != "GET" {
		var buf bytes.Buffer
		if err := gob.NewEncoder(&buf).Encode(value); err != nil {
			log.Println("Error occured while encoding ", key, " data =", err)
			return "", err
		}
		payloadValue = buf.String()
	}

	url := "http://127.0.0.1:12380/" + key
	var payload *strings.Reader
	payload = nil
	if value != nil {
		payload = strings.NewReader(payloadValue)
	}

	req, _ := http.NewRequest(method, url, payload)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error received from Raft =", err)
		return "", err
	}

	var data []byte
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error occured while decoding response from Raft =", err)
		return "", err
	}

	log.Println("data received from Raft after calling ", method, " method =", string(data))

	return string(data), nil
}

func NewToken() *authpb.AuthToken {
	return &authpb.AuthToken{
		Token: make(map[string]int32),
	}
}

func (*Server) GetUserIdFromToken(ctx context.Context, tokName *authpb.AuthTokenName) (*authpb.AuthTokenValue, error) {
	token := &authpb.AuthTokenValue{}
	var db TokenDB
	tokenDB, err := GetTokenDB(db.t)
	if err != nil {
		return nil, err
	}

	db.t = tokenDB

	log.Println("tokenDB after getting value from Raft Storage =", db.t)

	userId, ok := db.t.Token[tokName.TokenName]
	if ok {
		token.TokenValue = userId
		return token, nil
	} else {
		return token, errors.New("No token found or token expired!!")
	}
}

func (s *Server) AddToken(ctx context.Context, userId *authpb.UserId) (*authpb.AuthTokenName, error) {
	log.Println("Add Token called")
	var db TokenDB
	tok, err := s.GenerateToken(ctx, &authpb.InitToken{})
	if err != nil {
		log.Println("Error occured while adding token for userId", userId.User, " = ", err)
		return &authpb.AuthTokenName{TokenName: ""}, err
	}

	tokenDB, err := GetTokenDB(db.t)
	if err != nil {
		return nil, err
	}

	db.t = tokenDB

	db.t.Token[tok.TokenName] = userId.User

	_, err = InteractWithRaftStorage("PUT", "tokenDB", db.t)
	if err != nil {
		log.Println("Error occured while storing token data in Raft =", err)
		panic(err)
	}

	log.Println("Token generated for userId ", userId.User, " = ", tok.TokenName)
	log.Println("TokenDB = ", db.t.Token)

	return tok, err
}

func (*Server) UnsetToken(ctx context.Context, tokName *authpb.AuthTokenName) (*authpb.Status, error) {
	status := &authpb.Status{}
	var db TokenDB
	tokenDB, err := GetTokenDB(db.t)
	if err != nil {
		return nil, err
	}

	db.t = tokenDB
	delete(db.t.Token, tokName.TokenName)
	_, err = InteractWithRaftStorage("PUT", "tokenDB", db.t)
	if err != nil {
		log.Println("Error occured while storing token data in Raft =", err)
		panic(err)
	}
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
