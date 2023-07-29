package growl

// Notify sends a desktop notification
type Growl struct {
	appName string
	appIcon string
}

func (g *Growl) Growl(msg string) {
	g.growl(msg)
}

func New(appName string, appIcon string) *Growl {
	return new(appName, appIcon)
}

func (g *Growl) SetAppName(appName string) *Growl {
	return g.setAppName(appName)
}

func (g *Growl) SetAppIcon(appIcon string) *Growl {
	return g.setAppIcon(appIcon)
}
