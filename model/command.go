package model

import "time"

type Command struct {
	Duration time.Duration `json:"duration"`
}

func NewCommand() *Command {
	return &Command{
		1,
	}
}
