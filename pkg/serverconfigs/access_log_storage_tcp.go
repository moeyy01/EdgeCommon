// Copyright 2021 GoEdge CDN goedge.cdn@gmail.com. All rights reserved.

package serverconfigs

// AccessLogTCPStorageConfig TCP存储策略
type AccessLogTCPStorageConfig struct {
	Network string `yaml:"network" json:"network"` // tcp, unix
	Addr    string `yaml:"addr" json:"addr"`
}
