package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type ipInfo struct {
	IP   string `yaml:"ip"`
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}
type ipInfoList []*ipInfo

type configuration struct {
	Homepage string     `yaml:"homepage"`
	Title    string     `yaml:"title"`
	Timeout  int32      `yaml:"timeout"`
	IPList   ipInfoList `yaml:"ip_list"`
	ASNList  []struct {
		ASN  string `yaml:"asn"`
		Name string `yaml:"name"`
	} `yaml:"asn_list"`
}

func (infos ipInfoList) ListIP() []string {
	list := []string{}
	for _, info := range infos {
		list = append(list, info.IP)
	}
	return list
}

func (infos ipInfoList) ListName() []string {
	list := []string{}
	for _, info := range infos {
		list = append(list, info.Name)
	}
	return list
}

func (cfg *configuration) makeIPInfoList(ty string) ipInfoList {
	infos := ipInfoList{}
	for _, info := range cfg.IPList {
		if info.Type == ty {
			infos = append(infos, info)
		}
	}
	return infos
}

func (cfg *configuration) makeASNMap() map[string]string {
	set := make(map[string]string)
	for _, info := range cfg.ASNList {
		set[info.ASN] = info.Name
	}
	return set
}

func loadConfig(confFile string) (*configuration, error) {
	data, err := ioutil.ReadFile(confFile)
	if err != nil {
		return nil, err
	}

	cfg := configuration{}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
