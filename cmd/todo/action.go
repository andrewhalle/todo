package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/urfave/cli"

	"github.com/andrewhalle/todo/internal/common"
	"github.com/andrewhalle/todo/task"
)

func initialize(c *cli.Context) error {
	wd, _ := os.Getwd()
	err := initTodoDirectory(wd)
	common.CheckDie(err)
	return nil
}

func clean(c *cli.Context) error {
	wd, _ := os.Getwd()
	err := removeTodoDirectory(wd)
	common.CheckDie(err)
	return nil
}

func list(c *cli.Context) error {
	wd, _ := os.Getwd()
	dir := todoDirectoryPath(wd)
	tasks := task.FromDir(dir)

	alg := c.String("alg")
	if alg == "srtf" {
		tasks.SRTFSort()
	} else if alg == "fcfs" {
		tasks.FCFSSort()
	} else {
		return errors.New("unknown scheduling algorithm given")
	}

	for _, t := range tasks {
		fmt.Println("Name: ", t.Name)
		fmt.Println("Arrival Time: ", t.ArrivalTime)
		fmt.Println("Estimated Time: ", t.EstimatedTime)
		fmt.Println("Estimated Time Remaining: ", t.EstimatedTimeRemaining)
		fmt.Println("Time Spent: ", t.TimeSpent)
		fmt.Println("Priority: ", t.Priority)
		fmt.Println("")
	}
	return nil
	/*
		err := termbox.Init()
		common.CheckDie(err)
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
	*/
}

func add(c *cli.Context) error {
	wd, err := os.Getwd()
	common.CheckDie(err)
	dir := todoDirectoryPath(wd)
	t := task.FromUser()
	id, err := uuid.NewUUID()
	common.CheckDie(err)
	filename := id.String()
	if filename == "" {
		log.Fatal("uuid not valid")
	}
	t.Save(taskName(dir, filename))
	return nil
}
