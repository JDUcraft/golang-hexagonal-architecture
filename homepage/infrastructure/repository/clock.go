package repository

import (
	"time"
)

type Clock struct {
}

var Timer = time.Now

func (c Clock) GetTime() string {
	return Timer().Format("15:04:05")
}
