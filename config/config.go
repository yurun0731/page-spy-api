package config

import "io/fs"

type CrosConfig struct {
	AllowOrigins  []string `json:"allowOrigins"`
	AllowMethods  []string `json:"allowMethods"`
	AllowHeaders  []string `json:"allowHeaders"`
	ExposeHeaders []string `json:"exposeHeaders"`
}

type Config struct {
	MachineInfo *MachineInfo `json:"machineInfo"`
	Port        string       `json:"port"`
	CrosConfig  *CrosConfig  `json:"crosConfig"`
}

type Address struct {
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

type MachineInfo struct {
	MachineAddress map[string]*Address `json:"machineAddress"`
}
type StaticConfig struct {
	DirName string
	Files   fs.FS
}
