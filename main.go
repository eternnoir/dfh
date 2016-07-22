package main

import (
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	initApp(app)
	app.Run(os.Args)
}

func initApp(app *cli.App) {
	app.Name = "dfh"
	app.Usage = "Datetime base file helper. Easy to remove/move/find files by datetime."
	app.Version = "v0.1.0"
	InitCommands(app)
}
