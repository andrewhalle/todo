package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"fmt"
	"bufio"
	"strconv"
	"strings"

	//"github.com/nsf/termbox-go"
	"github.com/urfave/cli"
	"github.com/google/uuid"

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
	if err != nil {
		log.Fatal(err)
	}
}

/****************************
*      Helper functions     *
*****************************/

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

func check(e error) {
	if e != nil {
		log.Fatal("error: ", e)
	}
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
	/*
	err := termbox.Init()
	if err != nil {
		log.Fatal("error: ", err)
	}
	defer termbox.Close()
	w, h := termbox.Size()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for i := 0; i < w; i++ {
		termbox.SetCell(i, 0, '#', termbox.ColorDefault, termbox.ColorDefault)
	}
	for i := 1; i < h - 2; i++ {
		termbox.SetCell(0, i, '#', termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(w - 1, i, '#', termbox.ColorDefault, termbox.ColorDefault)
	}
	for i := 0; i < w; i++ {
		termbox.SetCell(i, h - 2, '#', termbox.ColorDefault, termbox.ColorDefault)
	}
	termbox.SetCell(0, h - 1, ':', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCursor(w / 2, h / 2)
	fmt.Println("hello world")
	termbox.Flush()
	time.Sleep(5 * time.Second)
	return nil
	*/
}

func add(c *cli.Context) error {
	wd, err := os.Getwd()
	check(err)
	dir := todoDirectoryPath(wd)
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Name: ")
	name, err := inputReader.ReadString('\n')
	check(err)
	fmt.Println("Time remaining: ")
	ttlStr, err := inputReader.ReadString('\n')
	check(err)
	ttl, err := strconv.ParseFloat(strings.TrimSpace(ttlStr), 64)
	check(err)
	fmt.Println("Priority: ")
	priorityStr, err := inputReader.ReadString('\n')
	check(err)
	priority, err := strconv.Atoi(strings.TrimSpace(priorityStr))
	check(err)
	t := task.Task{
		Name:           name,
		TimeToComplete: ttl,
		Priority:       priority,
	}
	id, err := uuid.NewUUID()
	check(err)
	filename := id.String()
	if filename == "" {
		log.Fatal("uuid not valid")
	}
	t.Save(taskName(dir, filename))
	return nil
}
