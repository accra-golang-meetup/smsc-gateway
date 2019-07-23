package main

type connectionManager struct {
	name string
}

func (connMgr *connectionManager) getName() string {
	return connMgr.name
}

func newConnectionManager(name string) *connectionManager {
	return &connectionManager{name: name}
}
