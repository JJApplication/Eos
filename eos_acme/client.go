/*
Create: 2022/8/24
Project: Eos
Github: https://github.com/landers1037
Copyright Renj
*/

// Package eos
package eos

import (
	"fmt"

	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/providers/http/webroot"
	"github.com/go-acme/lego/v4/registration"
)

type KeyType string

const (
	RSA2048 KeyType = "RSA-2048"
	RSA4096 KeyType = "RSA-4096"
	RSA8192 KeyType = "RSA-8192"
	EC256   KeyType = "ECDSA-256"
	EC384   KeyType = "ECDSA-384"
)

type ApiVersion string

const (
	Production ApiVersion = "Production"
	Sandbox    ApiVersion = "Sandbox"
)

// NewClient returns a new Lets Encrypt client
func NewClient(account *Account) (*lego.Client, error) {
	cfg := lego.NewConfig(account)
	switch account.KeyType {
	case RSA2048:
		cfg.Certificate.KeyType = certcrypto.RSA2048
		return lego.NewClient(cfg)
	case RSA4096:
		cfg.Certificate.KeyType = certcrypto.RSA4096
		return lego.NewClient(cfg)
	case RSA8192:
		cfg.Certificate.KeyType = certcrypto.RSA8192
		return lego.NewClient(cfg)
	case EC256:
		cfg.Certificate.KeyType = certcrypto.EC256
		return lego.NewClient(cfg)
	case EC384:
		cfg.Certificate.KeyType = certcrypto.EC384
		return lego.NewClient(cfg)
	default:
		return nil, fmt.Errorf("invalid private key type: %s", string(account.KeyType))
	}
}

// RegProvider 注册
func RegProvider(c *lego.Client) error {
	var err error
	ps, err := webroot.NewHTTPProvider(config.Challenge)
	if err != nil {
		return err
	}
	err = c.Challenge.SetHTTP01Provider(ps)
	if err != nil {
		return err
	}
	return err
}

func RegAccount(c *lego.Client) (*registration.Resource, error) {
	reg, err := c.Registration.Register(registration.RegisterOptions{TermsOfServiceAgreed: true})
	if err != nil {
		return reg, err
	}
	return reg, nil
}

func ObtainDomains(c *lego.Client, domains []string) (*certificate.Resource, error) {
	req := certificate.ObtainRequest{Domains: domains, Bundle: true}
	return c.Certificate.Obtain(req)
}
