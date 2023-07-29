package growl

func (g *growl) Growl(msg string) {
	g.growl(msg)
}

func New(appName string, appIcon string) *growl {
	return new(appName, appIcon)
}

func (g *growl) SetAppName(appName string) *growl {
	return g.setAppName(appName)
}

func (g *growl) SetAppIcon(appIcon string) *growl {
	return g.setAppIcon(appIcon)
}
