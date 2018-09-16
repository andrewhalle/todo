package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"errors"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/urfave/cli"
	"github.com/andrewhalle/todo/task"
)

func main() {
	app := cli.NewApp()
	app.Name = "todo"
	app.Usage = "decide what to do next using your favorite scheduling algorithm"
	app.Version = "0.0.1"
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
			Action: initialize,
		},
		{
			Name: "clean",
			Usage: "remove the .todo directory",
			Action: clean,
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

func getTermSize() (int, int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	output, _ := cmd.Output()
	stty_output := strings.Split(strings.TrimSpace(string(output)), " ")
	sizes := make([]int, 2)
	for i, size := range stty_output {
		sizes[i], _ = strconv.Atoi(size)
	}
	return sizes[0], sizes[1]
}

func clearTerm() {
	cmd := exec.Command("tput", "clear")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Run()
}

func swapToDir(path string) string {
	wd, _ := os.Getwd()
	os.Chdir(path)
	return wd
}

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
		oldWd := swapToDir(path)
		defer swapToDir(oldWd)
		os.Mkdir(".todo", 0755)
		return nil
	}
	return errors.New("the todo directory has already been initialized!")
}

func removeTodoDirectory(path string) error {
	if todoDirectoryExists(path) {
		oldWd := swapToDir(path)
		defer swapToDir(oldWd)
		os.Remove(".todo")
		return nil
	}
	return errors.New("the todo directory doesn't exist!")
}

func taskName(dir string, uuid string) string {
	return dir + string(filepath.Separator) + uuid + ".task"
}

/****************************
*      Action functions     *
*****************************/

func initialize(c *cli.Context) error {
	wd, _ := os.Getwd()
	err := initTodoDirectory(wd)
	if err != nil {
		return err
	}
	return nil
}

func clean(c *cli.Context) error {
	wd, _ := os.Getwd()
	err := removeTodoDirectory(wd)
	if err != nil {
		return err
	}
	return nil
}

func list(c *cli.Context) error {
/*
	wd, _ := os.Getwd()
	dir := todoDirectoryPath(wd)
	tasks := task.FromDir(dir)
	for _, t := range tasks {
		fmt.Println("Name: ", t.Name)
		fmt.Println("Time to Complete: ", t.TimeToComplete)
		fmt.Println("Priority: ", t.Priority)
		fmt.Println("")
	}
	return nil
*/
	clearTerm()
	rows, cols := getTermSize()
	fmt.Println(strings.Repeat("#", cols))
	for i := 0; i < rows - 3; i++ {
		fmt.Println("#" + strings.Repeat(" ", cols - 2) + "#")
	}
	fmt.Println(strings.Repeat("#", cols))
	return nil
}

func add(c *cli.Context) error {
	wd, _ := os.Getwd()
	dir := todoDirectoryPath(wd)
	t := task.Task{
		Name: "added task",
		TimeToComplete: 1,
		Priority: 1,
	}
	t.Save(taskName(dir, "test"))
	return nil
}
