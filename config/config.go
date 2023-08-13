package config

import (
	"fmt"

	xdl "github.com/mt1976/appFrame/dataloader"
)

func init() {
	//Config = *xdl.New("config", "", "")
}

func New(name, path string) xdl.Payload {
	fmt.Println("init config")
	return *xdl.New(name, "", path)
}

func Debug(c xdl.Payload) {
	c.Debug()
}
