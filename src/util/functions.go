package util

import (
	"time"
	"os"
	"os/exec"
)

func Clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func GetFormattedTime() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05.00")
}