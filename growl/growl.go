package growl

import (
	xnote "github.com/gen2brain/beeep"
	xstr "github.com/mt1976/appFrame/strings"
)

// Notify sends a desktop notification
type growl struct {
	appName string
	appIcon string
}

func (g *growl) growl(msg string) {
	xnote.Notify(g.appName, msg, g.appIcon)
}

func new(appName string, appIcon string) *growl {

	if appName == "" {
		appName = xstr.SBracket("Application Name")
	}

	return &growl{appName: appName, appIcon: appIcon}
}

func (g *growl) setAppName(appName string) *growl {
	g.appName = appName
	return g
}

func (g *growl) setAppIcon(appIcon string) *growl {
	g.appIcon = appIcon
	return g
}
