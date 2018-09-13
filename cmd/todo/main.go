package main

import (
	"fmt"
	"log"
	"os"
	"errors"
	"path/filepath"

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
		{
			Name: "init",
			Usage: "initialize empty .todo directory",
			Action: todoInit,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

/****************************
*      Helper functions     *
*****************************/

func todoDirectoryPath(path string) string {
	return path + string(filepath.Separator) + ".todo"
}

func todoDirectoryExists(path string) bool {
	_, err := os.Stat(todoDirectoryPath(path))
	if err != nil {
		return false
	} else {
		return true
	}
}

func initTodoDirectory(path string) error {
	if !todoDirectoryExists(path) {
		os.Mkdir(todoDirectoryPath(path), 0600)
		return nil
	}
	return errors.New("the todo directory has already been initialized!")
}

/****************************
*      Action functions     *
*****************************/

func todoInit(c *cli.Context) error {
	wd, _ := os.Getwd()
	err := initTodoDirectory(wd)
	if err != nil {
		return err
	}
	return nil
}

func list(c *cli.Context) error {
	fmt.Println("You're all done!")
	return nil
}

func add(c *cli.Context) error {
	fmt.Println("Can't do that right now.")
	return nil
}
