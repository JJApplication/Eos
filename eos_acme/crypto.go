/*
Create: 2022/8/24
Project: Eos
Github: https://github.com/landers1037
Copyright Renj
*/

// Package eos
package eos

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

// 默认为ecdsa
func generatePrivateKey(keyType string) (*ecdsa.PrivateKey, error) {
	pk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	return pk, err
}
