package environment

import (
	xsys "github.com/mt1976/appFrame/system"
)

// PATH: application\environment.go
// Language: go

// Contains Basic Application environment Information
type environment struct {
	dockerURI              string `mapstructure:"dockerURI"`
	dockerPORT             string `mapstructure:"dockerPORT"`
	dockerPROTOCOL         string `mapstructure:"dockerPROTOCOL"`
	appName                string `mapstructure:"appName"`
	appVersion             string `mapstructure:"appVersion"`
	appURI                 string `mapstructure:"appURI"`
	appPORT                string `mapstructure:"appPORT"`
	appPROTOCOL            string `mapstructure:"appPROTOCOL"`
	appGlyph               string `mapstructure:"appGLYPH"`
	appGlyphColor          string `mapstructure:"appGLYPHCOLOR"`
	userName               string `mapstructure:"userName"`
	userPassword           string `mapstructure:"userPassword"`
	userSecret             string `mapstructure:"userSecret"`
	appTemplate            string `mapstructure:"appTemplate"`
	additionalServices     bool   `mapstructure:"additionalServices"`
	additionalServicesList []string
}

// Contains Basic Application Environment Information
var Application environment

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
	return Application.AppName()
}

// Return HostName
func HostName() string {
	return xsys.Get().Hostname
}

// Return Application HTTP Port
func ApplicationHTTPPort() string {
	return Application.appPORT
}

func (obj environment) DockerURI() string {
	return obj.dockerURI
}

func (obj environment) DockerPORT() string {
	return obj.dockerPORT
}

func (obj environment) DockerPROTOCOL() string {
	return obj.dockerPROTOCOL
}

func (obj environment) AppName() string {
	return obj.appName
}

func (obj environment) AppVersion() string {
	return obj.appVersion
}

func (obj environment) AppURI() string {
	return obj.appURI
}

func (obj environment) AppPORT() string {
	return obj.appPORT
}

func (obj environment) AppPROTOCOL() string {
	return obj.appPROTOCOL
}

func (obj environment) AppGlyph() string {
	return obj.appGlyph
}

func (obj environment) AppGlyphColor() string {
	return obj.appGlyphColor
}

func (obj environment) UserName() string {
	return obj.userName
}

func (obj environment) UserPassword() string {
	return obj.userPassword
}

func (obj environment) UserSecret() string {
	return obj.userSecret
}

func (obj environment) AppTemplate() string {
	return obj.appTemplate
}

func (obj environment) AdditionalServices() bool {
	return obj.additionalServices
}

func (obj environment) AdditionalServicesList() []string {
	return obj.additionalServicesList
}
