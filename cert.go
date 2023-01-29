/*
Create: 2022/8/24
Project: Eos
Github: https://github.com/landers1037
Copyright Renj
*/

// Package main
package main

import (
	eos "github.com/JJApplication/eos/eos_acme"
)

// 创建申请新的ssl
func obtainCert(cf eos.EosConfig) {
	// new account
	account, err := eos.NewAccount(cf.Email, eos.RSA2048)
	if err != nil {
		eos.GetLogger().ErrorF("init account error: %s", err.Error())
		return
	}
	// new client
	client, err := eos.NewClient(account)
	if err != nil {
		eos.GetLogger().ErrorF("init client error: %s", err.Error())
		return
	}

	// reg
	err = eos.RegProvider(client)
	if err != nil {
		eos.GetLogger().ErrorF("init reg web error: %s", err.Error())
		return
	}

	// reg account
	reg, err := eos.RegAccount(client)
	if err != nil {
		eos.GetLogger().ErrorF("init reg account error: %s", err.Error())
		return
	}

	account.AddRegistration(reg)
	res, err := eos.ObtainDomains(client, cf.Domain)
	if err != nil {
		eos.GetLogger().ErrorF("obtain cert error: %s", err.Error())
		return
	}
	err = eos.SaveCert(res)
	if err != nil {
		eos.GetLogger().ErrorF("save cert error: %s", err.Error())
		return
	}
}
