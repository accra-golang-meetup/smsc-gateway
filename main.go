package main

import "fmt"

func main() {
	cfgMgr := newConfigManager()
	connMgr := newConnectionManager("foo")
	fmt.Printf("cfg = %+v\n", cfgMgr.smsc)
	fmt.Printf("name = %+v\n", connMgr.name)
}
