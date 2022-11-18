package environment

import (
	xsys "github.com/mt1976/appFrame/system"
)

// PATH: application\environment.go
// Language: go

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
	return Application.appName
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

func (obj environment) DockerPort() string {
	return obj.dockerPORT
}

func (obj environment) DockerProtocol() string {
	return obj.dockerPROTOCOL
}

func (obj environment) Name() string {
	return obj.appName
}

func (obj environment) Version() string {
	return obj.appVersion
}

func (obj environment) URI() string {
	return obj.appURI
}

func Port() string {
	return Application.appPORT
}

func (obj environment) Protocol() string {
	return obj.appPROTOCOL
}

func (obj environment) Glyph() string {
	return obj.appGlyph
}

func (obj environment) GlyphColor() string {
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

func (obj environment) Template() string {
	return obj.appTemplate
}

func (obj environment) AdditionalServices() bool {
	return obj.additionalServices
}

func (obj environment) AdditionalServicesList() []string {
	return obj.additionalServicesList
}
