package model

import (
	"time"
)

type ResultMgo struct {
	ScheduleTaskName string `json:"task_name" bson:"task_name"`
	Result string `json:"result" bson:"result"`
	Time time.Time `json:"time" bson:"time"`
}

