package models

import (
	"errors"
	//"log"
	"math/rand"
)

type Token map[string]int

func NewToken() Token {
	return make(Token)
}

func (t Token) GetUserIdFromToken(token string) (int, error) {
	userId, ok := t[token]
	if ok {
		return userId, nil
	} else {
		return -1, errors.New("No token found or token expired!!")
	}
}

func (t Token) AddToken(userId int) string {
	token := GenerateToken()
	t[token] = userId
	return token
}

func (t Token) UnsetToken(token string) {
	delete(t, token)
}

//Generate 16 digit random numbers which is used as a token
func GenerateToken() string {
	letterBytes := "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, 16)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
