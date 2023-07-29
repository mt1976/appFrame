package growl

import (
	xnote "github.com/gen2brain/beeep"
	xstr "github.com/mt1976/appFrame/strings"
)

func (g *Growl) growl(msg string) {
	xnote.Notify(g.appName, msg, g.appIcon)
}

func new(appName string, appIcon string) *Growl {

	if appName == "" {
		appName = xstr.SBracket("Application Name")
	}

	return &Growl{appName: appName, appIcon: appIcon}
}

func (g *Growl) setAppName(appName string) *Growl {
	g.appName = appName
	return g
}

func (g *Growl) setAppIcon(appIcon string) *Growl {
	g.appIcon = appIcon
	return g
}
