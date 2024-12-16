package models

import "time"

type Task struct {
	Id        string
	Time      time.Time
	CmdString string
	IsFinish  bool
	Out       string
}
