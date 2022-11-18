package environment

import (
	xsys "github.com/mt1976/appFrame/system"
)

// PATH: application\environment.go
// Language: go

// Contains Basic Application Environment Information
type Environment struct {
	DockerURI              string `mapstructure:"dockerURI"`
	DockerPORT             string `mapstructure:"dockerPORT"`
	DockerPROTOCOL         string `mapstructure:"dockerPROTOCOL"`
	AppName                string `mapstructure:"appName"`
	AppVersion             string `mapstructure:"appVersion"`
	AppURI                 string `mapstructure:"appURI"`
	AppPORT                string `mapstructure:"appPORT"`
	AppPROTOCOL            string `mapstructure:"appPROTOCOL"`
	AppGlyph               string `mapstructure:"appGLYPH"`
	AppGlyphColor          string `mapstructure:"appGLYPHCOLOR"`
	UserName               string `mapstructure:"userName"`
	UserPassword           string `mapstructure:"userPassword"`
	UserSecret             string `mapstructure:"userSecret"`
	AppTemplate            string `mapstructure:"appTemplate"`
	AdditionalServices     bool   `mapstructure:"additionalServices"`
	AdditionalServicesList []string
}

// Contains Basic Application Environment Information
var Application Environment

// Contains Overrides for Application Environment Information
var Config map[string]string

// Contains Overrides for Application Environment Information
var Overrides map[string]string

// Contains Extra Information to be added to Application Environment Information
var Extras map[string]string

// Contains Nothing
var DUMMY map[string]string

func Refresh() {
	refresh()
}

func Debug() {
	debug()
}

func GetConfig(orig string, inName string, what string) string {
	return getConfig(orig, inName, what)
}

func GetOverride(orig string, inName string, what string) string {
	return getOverride(orig, inName, what)
}

func GetExtra(orig string, inName string, what string) string {
	return getExtra(orig, inName, what)
}

// Return AppName
func ApplicationName() string {
	return Application.AppName
}

// Return HostName
func HostName() string {
	return xsys.Get().Hostname
}

// Return Application HTTP Port
func ApplicationHTTPPort() string {
	return Application.AppPORT
}
