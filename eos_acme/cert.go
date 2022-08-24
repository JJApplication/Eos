/*
Create: 2022/8/24
Project: Eos
Github: https://github.com/landers1037
Copyright Renj
*/

// Package eos
package eos

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/go-acme/lego/v4/certificate"
)

const (
	PrivateKey  = "private.key"     // 私钥
	Certificate = "certificate.pem" // 证书
	CA          = "ca.crt"          // CA
	CSR         = "csr"             // 全量csr
)

// SaveCert 保存cert文件
// 默认会保存到$CertRoot下
func SaveCert(cert *certificate.Resource) error {
	root := config.CertRoot
	if root == "" {
		root = "."
	}
	if _, err := os.Stat(root); os.IsNotExist(err) {
		os.MkdirAll(root, 0755)
	}
	// write priv
	err := writeFile(PrivateKey, cert.PrivateKey)
	if err != nil {
		logger.ErrorF("write %s error: %s", PrivateKey)
		return err
	}
	err = writeFile(Certificate, cert.Certificate)
	if err != nil {
		logger.ErrorF("write %s error: %s", Certificate)
		return err
	}
	err = writeFile(CA, cert.IssuerCertificate)
	if err != nil {
		logger.ErrorF("write %s error: %s", CA)
		return err
	}
	err = writeFile(CSR, cert.CSR)
	if err != nil {
		logger.ErrorF("write %s error: %s", CSR)
		return err
	}
	return nil
}

func writeFile(f string, data []byte) error {
	fp := filepath.Join(config.CertRoot, f)
	return ioutil.WriteFile(fp, data, 0644)
}
