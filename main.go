package main

import (
	"fmt"
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
	app.Action = func(c *cli.Context) error {
		fmt.Println("boom! I say!")
		return nil
	}

	InitCommands(app)
}
