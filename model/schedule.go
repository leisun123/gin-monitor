package model

import (
	"gin-monitor/lib/cron"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type Schedule struct {
	gorm.Model
	ScheduleID string `json:"schedule_id" bson:"schedule_id" gorm:"primary_key"`
	TaskId      string `json:"task_id" bson:"task_id"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	Cron        string        `json:"cron" bson:"cron"`
	EntryId     cron.EntryID  `json:"entry_id" bson:"entry_id"`
	Status      string        `json:"status" bson:"status"`
	Enabled     bool          `json:"enabled" bson:"enabled"`
}

func (sch *Schedule) Save() error {
	db, err := getDB()
	if err != nil{
		return err
	}
	defer db.Close()
	result := db.Model(&Schedule{}).Where("schedule_id=?", sch.ScheduleID).Update(&sch)
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (sch *Schedule) Delete() error {
	db, err := getDB()
	if err != nil{
		return err
	}
	defer db.Close()
	result := db.Model(&Schedule{}).Where("schedule_id=?", sch.ScheduleID).Delete(&sch)
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// 查询所有的任务
func GetScheduleList() ([]Schedule, error) {
	db, err := getDB()
	if err != nil{
		return nil, err
	}
	defer db.Close()
	var schedule []Schedule
	result := db.Model(&Schedule{}).Find(&schedule)
	if err := result.Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return schedule,nil
}

// 通过ID查询定时任务
func GetSchedule(id string) (*Schedule, error) {
	db, err := getDB()
	if err != nil{
		return nil, err
	}
	defer db.Close()
	var schedule Schedule
	result := db.Model(&Schedule{}).Where("schedule_id=?", id).Find(&schedule)
	if err := result.Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return &schedule, nil
}

// 通过ID更新定时任务
func UpdateSchedule(id string, item ScheduleDto) error {
	db, err := getDB()
	if err != nil{
		return err
	}
	defer db.Close()
	schedule := TransScheduleDto(item)
	result := db.Model(&Schedule{}).Where("schedule_id=?", id).Update(&schedule)
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// 添加新定时任务
func AddSchedule(item ScheduleDto) error {
	db, err := getDB()
	if err != nil{
		return err
	}
	defer db.Close()
	sch := TransScheduleDto(item)
	sch.ScheduleID = uuid.NewV4().String()
	result := db.Model(&Schedule{}).Create(&sch)
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// 通过ID删除定时任务
func RemoveSchedule(id string) error {
	db, err := getDB()
	if err != nil{
		return err
	}
	defer db.Close()
	result := db.Model(&Schedule{}).Where("schedule_id=?", id).Delete(&Schedule{})
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// 获取所有的定时任务总数
func GetScheduleCount() (int, error) {
	db, err := getDB()
	if err != nil{
		return 0, err
	}
	defer db.Close()
	var count int
	result := db.Model(&Schedule{}).Find(&Schedule{}).Count(&count)
	if err := result.Error; err != nil {
		return 0,errors.WithStack(err)
	}

	return count, nil
}