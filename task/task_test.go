package task

import (
	"testing"
	"math/rand"
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
	start := Task{
		Name: "First Task",
		TimeToComplete: 1.0,
		Priority: 1,
	}
	start.Save(filename)
	loaded := Load(filename)
	if loaded.Name != "First Task" || loaded.TimeToComplete != 1.0 || loaded.Priority != 1 {
		t.Fail()
	}
}

