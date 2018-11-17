package models

import (
	"errors"
	"math/rand"
)

type Token map[int]int

func NewToken() Token {
	return make(Token)
}

func (t Token) GetToken(userId int) (int, error) {
	accessToken, ok := t[userId]
	if ok {
		return accessToken, nil
	} else {
		return -1, errors.New("No token found or token expired!!")
	}
}

func (t Token) AddToken(userId int) int {
	//t[userId] = GenerateToken()
	token := GenerateToken()
	t[token] = userId
	return token
}

//Generate 16 digit random numbers which is used as a token
func GenerateToken() int {
	return 10000000000000000 + rand.Intn(9999999999999999-10000000000000000)
}
