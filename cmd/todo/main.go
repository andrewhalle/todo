package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "todo"
	app.Usage = "decide what to do next using your favorite scheduling algorithm"
	app.Commands = []cli.Command{
		{
			Name:   "list",
			Usage:  "list all your tasks",
			Action: list,
		},
		{
			Name:   "add",
			Usage:  "add a task",
			Action: add,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

/****************************
*      Action functions     *
*****************************/

func list(c *cli.Context) error {
	fmt.Println("You're all done!")
	return nil
}

func add(c *cli.Context) error {
	fmt.Println("Can't do that right now.")
	return nil
}
