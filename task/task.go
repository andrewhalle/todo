package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/andrewhalle/todo/common"
)

type Task struct {
	Name           string
	TimeToComplete float64
	Priority       int
}

type Tasks []*Task

func (t *Task) Save(filename string) {
	b, err := json.Marshal(*t)
	common.CheckDie(err)
	ioutil.WriteFile(filename, b, 0644)
}

func Load(filename string) *Task {
	b, err := ioutil.ReadFile(filename)
	common.CheckDie(err)
	var t Task
	err = json.Unmarshal(b, &t)
	common.CheckDie(err)
	return &t
}

func FromDir(dir string) Tasks {
	dirFile, err := os.Open(dir)
	common.CheckDie(err)
	taskFilenames, err := dirFile.Readdirnames(-1)
	tasks := make([]*Task, 0, len(taskFilenames))
	for _, filename := range taskFilenames {
		tasks = append(tasks, Load(dir+string(filepath.Separator)+filename))
	}
	return tasks
}

func FromUser() *Task {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Name: ")
	name, err := inputReader.ReadString('\n')
	name = strings.TrimSpace(name)
	common.CheckDie(err)
	fmt.Print("Time remaining: ")
	ttlStr, err := inputReader.ReadString('\n')
	common.CheckDie(err)
	ttl, err := strconv.ParseFloat(strings.TrimSpace(ttlStr), 64)
	common.CheckDie(err)
	fmt.Print("Priority: ")
	priorityStr, err := inputReader.ReadString('\n')
	common.CheckDie(err)
	priority, err := strconv.Atoi(strings.TrimSpace(priorityStr))
	t := Task{
		Name:           name,
		TimeToComplete: ttl,
		Priority:       priority,
	}
	return &t
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
