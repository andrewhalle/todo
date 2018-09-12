package task

import (
	"io/ioutil"
	"encoding/json"
	"log"
)

type Task struct {
	Name string
	TimeToComplete float64
	Priority int
}

func check(e error) {
	if e != nil {
		log.Fatal("error: ", e)
	}
}

func (t *Task) Save(filename string) {
	b, err := json.Marshal(*t)
	check(err)
	ioutil.WriteFile(filename, b, 0644)
}

func Load(filename string) *Task {
	b, err := ioutil.ReadFile(filename)
	check(err)
	var t Task
	err = json.Unmarshal(b, &t)
	check(err)
	return &t
}

