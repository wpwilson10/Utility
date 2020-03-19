package setup

// ApplicationName is the currently running program
var ApplicationName string

// Application sets up global variables
func Application(app string) {
	ApplicationName = app
}
