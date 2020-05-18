package model

import (
	"github.com/apex/log"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type Task struct {
	gorm.Model
	TaskID string `json:"task_id" bson:"task_id" gorm:"primary_key"`
	TaskName string `json:"task_name" bson:"task_name"`
	TaskAPI string `json:"task_api" bson:"task_api"`
	Params string `json:"params" bson:"params"`
	WaitDuration    float64       `json:"wait_duration" bson:"wait_duration"`
	RuntimeDuration float64       `json:"runtime_duration" bson:"runtime_duration"`
	TotalDuration   float64       `json:"total_duration" bson:"total_duration"`
}


// 获取全部任务
func GetTaskList() ([]Task, error) {
	db, err := getDB()
	if err != nil{
		return nil, err
	}
	defer db.Close()
	var tasks []Task
	result := db.Model(&Task{}).Find(&tasks)
	if err := result.Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return tasks,nil
}

// 获取任务数量
func GetTaskCount() (int, error) {
	db, err := getDB()
	if err != nil{
		return 0, err
	}
	defer db.Close()
	var count int
	result := db.Model(&Task{}).Find(&Task{}).Count(&count)
	if err := result.Error; err != nil {
		return 0,errors.WithStack(err)
	}

	return count, nil
}

// 通过ID获取Task
func GetTask(id string) (*Task, error) {
	db, err := getDB()
	if err != nil{
		return nil, err
	}
	defer db.Close()
	var task Task
	result := db.Model(&Task{}).Where("task_id=?", id).Find(&task)
	if err := result.Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &task,nil
}

// 添加Task
func AddTask(item TaskDto) error {
	db, err := getDB()
	if err != nil{
		return err
	}
	defer db.Close()
	task := TransTaskDto(item)
	task.TaskID = uuid.NewV4().String()
	result := db.Model(&Task{}).Create(&task)
	if err := result.Error; err != nil {
		log.Errorf(err.Error())
		return errors.WithStack(err)
	}
	return nil
}

// 删除Task
func RemoveTask(id string) error {
	db, err := getDB()
	if err != nil{
		return err
	}
	defer db.Close()
	result := db.Model(Task{}).Where("task_id=?", id).Delete(&Task{})
	if err := result.Error; err != nil {
		log.Errorf(err.Error())
		return errors.WithStack(err)
	}
	return nil
}

// 更新task
func UpdateTask(id string, item TaskDto) error {
	db, err := getDB()
	if err != nil{
		return err
	}
	defer db.Close()
	task := TransTaskDto(item)
	result := db.Model(&Task{}).Where("task_id=?", id).Update(&task)
	if err := result.Error; err != nil {
		log.Errorf(err.Error())
		return errors.WithStack(err)
	}
	return nil
}
