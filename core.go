/*
Create: 2022/8/24
Project: Eos
Github: https://github.com/landers1037
Copyright Renj
*/

// Package main
package main

import (
	"time"

	eos "github.com/JJApplication/eos/eos_acme"
	"github.com/JJApplication/fushin/log"
)

type EosCore struct {
	Name      string
	ReNewTime string
	cf        eos.EosConfig
	logger    *log.Logger
}

func (e *EosCore) Run() {
	autoReNew(e.cf, e.ReNewTime)
	e.preRun()
	e.mainLoop()
}

func (e *EosCore) mainLoop() {
	ticker := time.Tick(time.Second * 60)
	for range ticker {
		e.logger.InfoF("[%s] auto check success", e.Name)
	}
}

func (e *EosCore) preRun() {
	e.logger.InfoF("[%s] config loaded: %+v", e.Name, e.cf)
}
