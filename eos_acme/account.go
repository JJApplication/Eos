/*
Create: 2022/8/24
Project: Eos
Github: https://github.com/landers1037
Copyright Renj
*/

// Package eos
package eos

import (
	"crypto"
	"encoding/json"
	"errors"
	"io/ioutil"
	"path"

	"github.com/go-acme/lego/v4/registration"
)

type Account struct {
	Email        string                 `json:"email"`
	Registration *registration.Resource `json:"registrations"`
	KeyType      KeyType                `json:"keyType"`
	key          crypto.PrivateKey
	path         string
}

// NewAccount 创建账户
func NewAccount(email string, keyType KeyType) (*Account, error) {
	if email == "" {
		return nil, errors.New("email is empty")
	}
	var privKey crypto.PrivateKey
	privKey, err := generatePrivateKey("")
	if err != nil {
		logger.ErrorF("generate private key error: %s", err.Error())
		return nil, err
	}

	account := &Account{
		Email:        email,
		Registration: nil,
		KeyType:      keyType,
		key:          privKey,
		path:         "",
	}

	return account, nil
}

// Save the account to disk
func (a *Account) Save() error {
	jsonBytes, err := json.MarshalIndent(a, "", "\t")
	if err != nil {
		return err
	}
	accountFile := path.Join(a.path, "account.json")
	return ioutil.WriteFile(accountFile, jsonBytes, 0700)
}

/* Methods implementing the lego.User interface*/

// GetEmail returns the email address for the account
func (a *Account) GetEmail() string {
	return a.Email
}

// GetPrivateKey returns the private RSA account key.
func (a *Account) GetPrivateKey() crypto.PrivateKey {
	return a.key
}

// GetRegistration returns the server registration
func (a *Account) GetRegistration() *registration.Resource {
	return a.Registration
}

func (a *Account) AddRegistration(reg *registration.Resource) {
	a.Registration = reg
}
