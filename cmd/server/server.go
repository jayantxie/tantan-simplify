package main

func main() {
	// parse command line args, not useful
	parseFlags()

	// show version of the project
	showVersion()

	// pre config dependency plugin
	setup()

	// start up HTTP Server
	startServer()
}
