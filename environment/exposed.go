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

// The Refresh function calls the refresh function.
func Refresh() {
	refresh()
}

// The function "Debug" calls the function "debug".
func Debug() {
	debug()
}

// The function GetConfig returns a configuration value based on the input parameters.
func GetConfig(orig string, inName string, what string) string {
	return getConfig(orig, inName, what)
}

// The function GetOverride returns the override value for a given input name and what parameter.
func GetOverride(orig string, inName string, what string) string {
	return getOverride(orig, inName, what)
}

// The function GetExtra returns a string value based on the input parameters orig, inName, and what.
func GetExtra(orig string, inName string, what string) string {
	return getExtra(orig, inName, what)
}

// Return AppName
// The function returns the name of the application.
func ApplicationName() string {
	return Application.AppName
}

// Return HostName
// The function returns the hostname of the system.
func HostName() string {
	return xsys.Get().Hostname
}

// Return Application HTTP Port
func ApplicationHTTPPort() string {
	return Application.AppPORT
}

// The function DockerURI returns the Docker URI of the application.
func DockerURI() string {
	return Application.DockerURI
}

// The function DockerPort returns the Docker port of the application.
func DockerPort() string {
	return Application.DockerPORT
}

// The function returns the Docker protocol used by the application.
func DockerProtocol() string {
	return Application.DockerPROTOCOL
}

// The Name function returns the name of the application.
func Name() string {
	return Application.AppName
}

// The Version function returns the application version.
func Version() string {
	return Application.AppVersion
}

// The URI function returns the URI of the application.
func URI() string {
	return Application.AppURI
}

// The function returns the value of the AppPORT variable in the Application package.
func Port() string {
	return Application.AppPORT
}

// The function returns the value of the AppPROTOCOL constant from the Application package.
func Protocol() string {
	return Application.AppPROTOCOL
}

// The function "Glyph" returns the application's glyph.
func Glyph() string {
	return Application.AppGlyph
}

// The function returns the application's glyph color.
func GlyphColor() string {
	return Application.AppGlyphColor
}

// The function returns the username of the application.
func UserName() string {
	return Application.UserName
}

// The function returns the user's password from the application.
func UserPassword() string {
	return Application.UserPassword
}

// The function UserSecret returns the user secret from the Application.
func UserSecret() string {
	return Application.UserSecret
}

// The Template function returns the application's template.
func Template() string {
	return Application.AppTemplate
}

// The function AdditionalServices returns a boolean value indicating whether additional services are
// available in the application.
func AdditionalServices() bool {
	return Application.AdditionalServices
}

// The function AdditionalServicesList returns a list of additional services.
func AdditionalServicesList() []string {
	return Application.AdditionalServicesList
}
