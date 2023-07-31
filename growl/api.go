package growl

// Notify sends a desktop notification
// The Growl type represents a notification with an app name and app icon.
// @property {string} appName - The appName property is a string that represents the name of the
// application.
// @property {string} appIcon - The `appIcon` property is a string that represents the icon of the
// application. It can be a file path or a URL pointing to the location of the icon image.
type Growl struct {
	appName string
	appIcon string
}

// The `func (g *Growl) Growl(msg string)` function is a method of the `Growl` struct. It takes a `msg`
// parameter of type string and calls the `growl` method of the `Growl` struct with the `msg`
// parameter. This method is responsible for sending a desktop notification.
func (g *Growl) Growl(msg string) {
	g.growl(msg)
}

// The New function creates a new instance of the Growl struct with the specified app name and icon.
func New(appName string, appIcon string) *Growl {
	return new(appName, appIcon)
}

// The `func (g *Growl) SetAppName(appName string) *Growl` function is a method of the `Growl` struct.
// It takes a `appName` parameter of type string and sets the `appName` property of the `Growl` struct
// to the specified value. It then returns a pointer to the modified `Growl` struct.
func (g *Growl) SetAppName(appName string) *Growl {
	return g.setAppName(appName)
}

// The `func (g *Growl) SetAppIcon(appIcon string) *Growl` function is a method of the `Growl` struct.
// It takes a `appIcon` parameter of type string and sets the `appIcon` property of the `Growl` struct
// to the specified value. It then returns a pointer to the modified `Growl` struct. This function
// allows you to update the app icon of the Growl notification.
func (g *Growl) SetAppIcon(appIcon string) *Growl {
	return g.setAppIcon(appIcon)
}
