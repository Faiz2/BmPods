package BmConfig

import "github.com/alfredyang1986/blackmirror/bmconfighandle"

type BmRouterConfig struct {
	Host string
	Port string
	TmpDir string
}

func (br *BmRouterConfig) GenerateConfig() {
	//TODO: 配置文件路径 待 用脚本指定dev路径和deploy路径
	configPath := "Resource/routerconfig.json"
	profileItems := bmconfig.BMGetConfigMap(configPath)

	br.Host = profileItems["Host"].(string)
	br.Port = profileItems["Port"].(string)
	br.TmpDir = profileItems["TmpDir"].(string)
}