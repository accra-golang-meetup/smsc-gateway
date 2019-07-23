package main

type configManager struct {
	smsc string
}

func (cfgMgr *configManager) getConfig() string {
	return cfgMgr.smsc
}

func newConfigManager() *configManager {
	return &configManager{smsc: "foo-smsc"}
}
