package main

import (
	"errors"
	"fmt"
	"github.com/eternnoir/dfh/lib"
	"github.com/urfave/cli"
	"os"
	"sync"
	"time"
)

func InitCommands(app *cli.App) {
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "timetemplate",
			Value: "2006-01-02T15:04:05Z07:00",
			Usage: "Parse time template.",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "remove",
			Aliases: []string{"rm"},
			Usage:   "remove files.",
			Flags:   getFlags(),
			Action:  removeAction,
		},
		{
			Name:    "find",
			Aliases: []string{"f"},
			Usage:   "find files.",
			Flags:   getFlags(),
			Action:  findAction},
	}
}

func getFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "outtime, o",
			Value: "",
			Usage: "How many time duration ago? eg. 24h",
		},
		cli.StringFlag{
			Name:  "basetime, b",
			Value: "",
			Usage: "Datetime base line.",
		},
		cli.BoolFlag{
			Name:  "recursive, r",
			Usage: "Use recursive.",
		},
		cli.BoolFlag{
			Name:  "include, i",
			Usage: "Include time range. Default exclude.",
		},
		cli.BoolFlag{
			Name:  "force, f",
			Usage: "Force, without asking.",
		},
		cli.BoolFlag{
			Name:  "dir, d",
			Usage: "Search directory.",
		},
	}
}

type Params struct {
	StartTime time.Time
	EndTime   time.Time
	Recursive bool
	Force     bool
	Include   bool
	Dir       bool
	Path      string
}

func ParseFlags(c *cli.Context) (Params, error) {
	timetemplateFlag := c.GlobalString("timetemplate")
	outtimeFlag := c.String("outtime")
	basetimeFlag := c.String("basetime")
	if c.NArg() < 1 {
		return Params{}, errors.New("Please give path.")
	}

	baseTime, err := ParseBaseTime(basetimeFlag, timetemplateFlag)
	if err != nil {
		fmt.Errorf("Parse BaseTime Fail. %s", err.Error())
		return Params{}, err
	}

	startTime, err := ParseStartTime(outtimeFlag, "", baseTime)
	if err != nil {
		fmt.Errorf("Parse StartTime Fail. %s", err.Error())
		return Params{}, err
	}

	return Params{
		StartTime: startTime,
		EndTime:   time.Unix(1<<63-62135596801, 999999999),
		Recursive: c.Bool("recursive"),
		Force:     c.Bool("force"),
		Include:   c.Bool("include"),
		Dir:       c.Bool("dir"),
		Path:      c.Args()[0],
	}, nil
}

func CreateFinder(param *Params) (lib.FileFInder, error) {
	return lib.NewFileFinder(param.StartTime, param.EndTime, param.Recursive, param.Include, param.Dir)
}

func findAction(c *cli.Context) error {
	params, err := ParseFlags(c)
	if err != nil {
		return cli.NewExitError(fmt.Sprintf("Parse params error. [Message] %s", err), 2)
	}
	finder, err := CreateFinder(&params)
	if err != nil {
		return cli.NewExitError(fmt.Sprintf("Create Finder error. [Message] %s", err), 3)
	}
	paths, dirs, err := finder.FindFiles(params.Path)
	if err != nil {
		return cli.NewExitError(fmt.Sprintf("Walk Path error. [Message] %s", err), 4)
	}

	printPath(dirs)
	printPath(paths)

	return nil

}

func removeAction(c *cli.Context) error {
	params, err := ParseFlags(c)
	if err != nil {
		return cli.NewExitError(fmt.Sprintf("Parse params error. [Message] %s", err), 2)
	}
	finder, err := CreateFinder(&params)
	if err != nil {
		return cli.NewExitError(fmt.Sprintf("Create Finder error. [Message] %s", err), 3)
	}
	paths, dirs, err := finder.FindFiles(params.Path)
	if err != nil {
		return cli.NewExitError(fmt.Sprintf("Walk Path error. [Message] %s", err), 4)
	}
	printPath(dirs)
	printPath(paths)
	if !params.Force {
		fmt.Print("Are you sure remove these file? (Y/N) ")
		var check string
		fmt.Scanln(&check)
		if check != "Y" {
			fmt.Println("Abort.")
			return nil
		}
	}
	removeFiles(paths)
	removeFiles(dirs)
	return nil
}

func removeFiles(paths []string) {
	var wg sync.WaitGroup
	for _, p := range paths {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			err := os.RemoveAll(path)
			if err != nil {
				fmt.Println(err)
			}
		}(p)
	}
	wg.Wait()
}

func printPath(paths []string) {
	for _, p := range paths {
		fmt.Println(p)
	}
}
