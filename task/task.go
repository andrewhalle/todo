package task

import (
	"io/ioutil"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"sort"
)

type Task struct {
	Name string
	TimeToComplete float64
	Priority int
}

type Tasks []*Task

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

func FromDir(dir string) Tasks {
	dirFile, err := os.Open(dir)
	check(err)
	taskFilenames, err := dirFile.Readdirnames(-1)
	tasks := make([]*Task, 0, len(taskFilenames))
	for _, filename := range taskFilenames {
		tasks = append(tasks, Load(dir + string(filepath.Separator) + filename))
	}
	return tasks
}

/****************************************************
*                    Sorting                        *
*****************************************************/

type byTimeToComplete []*Task

func (a byTimeToComplete) Len() int {
	return len(a)
}

func (a byTimeToComplete) Swap(i, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}

func (a byTimeToComplete) Less(i, j int) bool {
	return a[i].TimeToComplete < a[j].TimeToComplete
}

func (t Tasks) SJFSort() error {
	sort.Sort(byTimeToComplete(t))
	return nil
}

