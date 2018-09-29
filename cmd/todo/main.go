package main

import (
	"os"

	//"github.com/nsf/termbox-go"
	"github.com/urfave/cli"

	"github.com/andrewhalle/todo/internal/common"
)

func main() {
	app := cli.NewApp()
	app.Name = "todo"
	app.Usage = "decide what to do next using your favorite scheduling algorithm"
	app.Version = "0.1.0"
	app.Commands = []cli.Command{
		{
			Name:   "list",
			Usage:  "list all your tasks",
			Action: list,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "alg",
					Value: "fcfs",
					Usage: "scheduling algorithm to use",
				},
			},
		},
		{
			Name:   "add",
			Usage:  "add a task",
			Action: add,
		},
		{
			Name:   "init",
			Usage:  "initialize empty .todo directory",
			Action: initialize,
		},
		{
			Name:   "clean",
			Usage:  "remove the .todo directory",
			Action: clean,
		},
	}

	err := app.Run(os.Args)
	common.CheckDie(err)
}
