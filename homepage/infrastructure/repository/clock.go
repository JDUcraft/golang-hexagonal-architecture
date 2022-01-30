package repository

import (
	"time"
)

type Clock struct {
}

func (c Clock) GetTime() string {
	return time.Now().Format("15:04:05")
}
