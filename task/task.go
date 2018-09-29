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
	"time"

	"github.com/andrewhalle/todo/common"
)

type Task struct {
	Name                   string
	ArrivalTime            time.Time
	EstimatedTime          time.Duration
	EstimatedTimeRemaining time.Duration
	TimeSpent              time.Duration
	Priority               int
	StructVersion          int
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

	fmt.Print("Estimated time to complete: ")
	etStr, err := inputReader.ReadString('\n')
	common.CheckDie(err)
	et, err := time.ParseDuration(strings.TrimSpace(etStr))
	common.CheckDie(err)

	fmt.Print("Priority: ")
	priorityStr, err := inputReader.ReadString('\n')
	common.CheckDie(err)
	priority, err := strconv.Atoi(strings.TrimSpace(priorityStr))

	t := Task{
		Name:                   name,
		ArrivalTime:            time.Now(),
		EstimatedTime:          et,
		EstimatedTimeRemaining: et,
		TimeSpent:              time.Duration(0),
		Priority:               priority,
		StructVersion:          1,
	}
	return &t
}

/****************************************************
*                    Sorting                        *
*****************************************************/

type byEstimatedTimeRemaining []*Task

func (a byEstimatedTimeRemaining) Len() int {
	return len(a)
}

func (a byEstimatedTimeRemaining) Swap(i, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}

func (a byEstimatedTimeRemaining) Less(i, j int) bool {
	return a[i].EstimatedTimeRemaining < a[j].EstimatedTimeRemaining
}

func (t Tasks) SRTFSort() error {
	sort.Sort(byEstimatedTimeRemaining(t))
	return nil
}
