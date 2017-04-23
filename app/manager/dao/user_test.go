package dao

import (
	"testing"
)

const (
	email    = "569230199@qq.com"
	password = "12e32d23"
)

func TestRandToken(t *testing.T) {
	token := RandToken()
	if len(token) != 172 {
		t.Error(token)
	}
}

func TestCreateUser(t *testing.T) {
	err := CreateUser(email, password)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetUserByEmail(t *testing.T) {
	user, err := GetUserByEmail(email)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(user)
}

func TestGetUserByToken(t *testing.T) {
	user1, err := GetUserByEmail(email)
	if err != nil {
		t.Fatal(err)
	}
	user2, err := GetUserByToken(user1.Token)
	if err != nil {
		t.Fatal(err)
	}
	if user1.Email != user2.Email {
		t.Error(user2)
	}
	t.Log(user2)
}

func TestUserLogin(t *testing.T) {
	user, err := UserLogin(email, password)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(user)
	user, err = UserLogin(email, "error password")
	if err == nil {
		t.Fatal(user)
	}
}
