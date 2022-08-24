/*
Create: 2022/8/24
Project: Eos
Github: https://github.com/landers1037
Copyright Renj
*/

// Package eos
package eos

import (
	"github.com/JJApplication/fushin/log"
)

func init() {
	logger = log.Default("Eos")
	logger.Init()
}

var logger *log.Logger

func GetLogger() *log.Logger {
	return logger
}
