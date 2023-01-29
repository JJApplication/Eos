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

func main() {
	core := EosCore{
		Name:      "Eos",
		ReNewTime: "* * * * 1/3 ?",
		cf:        eos.GetConfig(),
		logger:    eos.GetLogger(),
	}

	core.Run()
}
