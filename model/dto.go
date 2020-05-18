package model

type ScheduleDto struct {
	TaskID string `json:"task_id" bson:"task_id"`
	ScheduleTaskName string `json:"schedule_task_name" bson:"schedule_task_name"`
	Description string        `json:"description" bson:"description"`
	Cron        string        `json:"cron" bson:"cron"`
}

type TaskDto struct {
	TaskName string `json:"task_name" bson:"task_name"`
	TaskAPI string	`json:"task_api" bson:"task_api"`
	Params string `json:"params" bson:"params"`
}

func TransScheduleDto(dto ScheduleDto) Schedule {
	return Schedule{
		TaskId:      dto.TaskID,
		Name:        dto.ScheduleTaskName,
		Description: dto.Description,
		Cron:        dto.Cron,
	}
}

func TransTaskDto(dto TaskDto) Task {
	return Task{
		TaskName: dto.TaskName,
		TaskAPI: dto.TaskAPI,
		Params: dto.Params,
	}
}