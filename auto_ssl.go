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
	"github.com/JJApplication/fushin/cron"
)

func autoReNew(cf eos.EosConfig, spec string) {
	g := cron.NewGroup(spec)
	id, err := g.AddFunc(func() {
		eos.GetLogger().WarnF("cronjob start to auto renew")
		obtainCert(cf)
		eos.GetLogger().WarnF("cronjob finished")
	})
	if err != nil {
		eos.GetLogger().ErrorF("cronjob init error: %s", err.Error())
	}
	eos.GetLogger().WarnF("cronjob init success: %d", id)
}
