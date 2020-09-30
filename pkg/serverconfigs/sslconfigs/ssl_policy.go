package sslconfigs

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
)

// TLS Version
type TLSVersion = string

// Cipher Suites
type TLSCipherSuite = string

// SSL配置
type SSLPolicy struct {
	Id   int64 `yaml:"id" json:"id"`     // ID
	IsOn bool  `yaml:"isOn" json:"isOn"` // 是否开启

	CertRefs       []*SSLCertRef     `yaml:"certRefs" json:"certRefs"`
	Certs          []*SSLCertConfig  `yaml:"certs" json:"certs"`
	ClientAuthType SSLClientAuthType `yaml:"clientAuthType" json:"clientAuthType"` // 客户端认证类型

	MinVersion   TLSVersion       `yaml:"minVersion" json:"minVersion"`     // 支持的最小版本
	CipherSuites []TLSCipherSuite `yaml:"cipherSuites" json:"cipherSuites"` // 加密算法套件

	HSTS         *HSTSConfig `yaml:"hsts2" json:"hsts"`                // HSTS配置，yaml之所以使用hsts2，是因为要和以前的版本分开
	HTTP2Enabled bool        `yaml:"http2Enabled" json:"http2Enabled"` // 是否启用HTTP2

	nameMapping map[string]*tls.Certificate // dnsName => cert

	minVersion   uint16
	cipherSuites []uint16

	clientCAPool *x509.CertPool
}

// 校验配置
func (this *SSLPolicy) Init() error {
	if len(this.Certs) == 0 {
		return errors.New("no certificates in https config")
	}

	for _, cert := range this.Certs {
		err := cert.Init()
		if err != nil {
			return err
		}
	}

	// min version
	this.convertMinVersion()

	// cipher suite categories
	this.initCipherSuites()

	// hsts
	if this.HSTS != nil {
		err := this.HSTS.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

// 取得最小版本
func (this *SSLPolicy) TLSMinVersion() uint16 {
	return this.minVersion
}

// 套件
func (this *SSLPolicy) TLSCipherSuites() []uint16 {
	return this.cipherSuites
}

// 校验是否匹配某个域名
func (this *SSLPolicy) MatchDomain(domain string) (cert *tls.Certificate, ok bool) {
	for _, cert := range this.Certs {
		if cert.MatchDomain(domain) {
			return cert.CertObject(), true
		}
	}
	return nil, false
}

// 取得第一个证书
func (this *SSLPolicy) FirstCert() *tls.Certificate {
	for _, cert := range this.Certs {
		return cert.CertObject()
	}
	return nil
}

// CA证书Pool，用于TLS对客户端进行认证
func (this *SSLPolicy) CAPool() *x509.CertPool {
	return this.clientCAPool
}
