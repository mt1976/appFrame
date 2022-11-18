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

func DockerURI() string {
	return Application.DockerURI
}

func DockerPort() string {
	return Application.DockerPORT
}

func DockerProtocol() string {
	return Application.DockerPROTOCOL
}

func Name() string {
	return Application.AppName
}

func Version() string {
	return Application.AppVersion
}

func URI() string {
	return Application.AppURI
}

func Port() string {
	return Application.AppPORT
}

func Protocol() string {
	return Application.AppPROTOCOL
}

func Glyph() string {
	return Application.AppGlyph
}

func GlyphColor() string {
	return Application.AppGlyphColor
}

func UserName() string {
	return Application.UserName
}

func UserPassword() string {
	return Application.UserPassword
}

func UserSecret() string {
	return Application.UserSecret
}

func Template() string {
	return Application.AppTemplate
}

func AdditionalServices() bool {
	return Application.AdditionalServices
}

func AdditionalServicesList() []string {
	return Application.AdditionalServicesList
}
