package task

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var chars []rune = []rune("abcdefghijklmnopqrstuvwxyz")

func randomFilename() string {
	b := make([]rune, 10)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return "/tmp/" + string(b) + ".task"
}

func TestTask(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	filename := randomFilename()
	now := time.Now()
	start := Task{
		Name:          "First Task",
		ArrivalTime:   now,
		EstimatedTime: time.Duration(0),
		Priority:      1,
	}
	start.Save(filename)
	loaded := Load(filename)
	if loaded.Name != "First Task" ||
		!loaded.ArrivalTime.Equal(now) ||
		loaded.EstimatedTime != time.Duration(0) ||
		loaded.Priority != 1 {

		fmt.Println(loaded)
		t.Fail()
	}
}
