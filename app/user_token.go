package main

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"
)

type UserToken struct {
	Id     int64
	UserId int64
	Token  string
}

func (model *UserToken) Issue(login string) (*UserToken, error) {
	model.Token = generate(login)

	err := model.insert()
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (model *UserToken) insert() error {
	_, err := App.Db.Exec(`
		INSERT INTO user_tokens(user_id, token)
		VALUES($1, $2)
	`, model.UserId, model.Token)

	return err
}

func generate(login string) string {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	salt := salt(10)

	hasher := md5.New()
	hasher.Write([]byte(login + timestamp + salt))
	return hex.EncodeToString(hasher.Sum(nil))
}

func salt(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
