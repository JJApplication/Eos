/*
Create: 2022/8/24
Project: Eos
Github: https://github.com/landers1037
Copyright Renj
*/

// Package eos
package eos

import (
	"strings"

	"github.com/JJApplication/fushin/utils/env"
)

type EosConfig struct {
	Email     string
	Domain    []string
	CertRoot  string
	Challenge string // 存储质询文件路径
}

var config EosConfig

// 初始化需要使用到的全部环境变量
func init() {
	envLoader := env.EnvLoader{}
	config = EosConfig{
		Email:     envLoader.Get("Email").Raw(),
		Domain:    strings.Fields(envLoader.Get("Domain").Raw()),
		CertRoot:  envLoader.Get("CertRoot").Raw(),
		Challenge: envLoader.Get("Challenge").Raw(),
	}
}

func GetConfig() EosConfig {
	return config
}
