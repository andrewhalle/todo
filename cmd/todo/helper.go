package main

import (
	"errors"
	"os"
	"path/filepath"
)

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
