package main

import (
	"gin-monitor/config"
	_ "gin-monitor/main/docs"
	"gin-monitor/model"
	"gin-monitor/routes"
	"gin-monitor/services"
	"github.com/apex/log"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"runtime/debug"
)

func main()  {
	// 初始化配置
	if err := config.InitConfig(""); err != nil {
		log.Error("init config error:" + err.Error())
		panic(err)
	}

	// 初始化Msql数据库
	if err := model.InitMysql(); err != nil {
		log.Error("init mysql error:" + err.Error())
		debug.PrintStack()
		panic(err)
	}
	log.Info("initialized mysql successfully")

	// 初始化定时任务
	if err := services.InitScheduler(); err != nil {
		log.Error("init scheduler error:" + err.Error())
		debug.PrintStack()
		panic(err)
	}
	log.Info("initialized schedule successfully")

	router := gin.Default()
	task := router.Group("/")
	{
		task.PUT("task", routes.NewTask)
		task.DELETE("task/:task_id", routes.RemoveTaskById)
		task.POST("task/:task_id", routes.UpdateTask)
		task.GET("task", routes.GetTaskList)
	}
	schedule := router.Group("/")
	// 定时任务
	{
		schedule.GET("schedules", routes.GetScheduleList)              // 定时任务列表
		schedule.GET("schedules/:id", routes.GetSchedule)              // 定时任务详情
		schedule.PUT("schedules", routes.PutSchedule)                  // 创建定时任务
		schedule.POST("schedules/:id", routes.PostSchedule)            // 修改定时任务
		schedule.DELETE("schedules/:id", routes.DeleteSchedule)        // 删除定时任务
		schedule.POST("schedules/:id/disable", routes.DisableSchedule) // 禁用定时任务
		schedule.POST("schedules/:id/enable", routes.EnableSchedule)   // 启用定时任务
	}
	//url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	_ = router.Run(":9090")
}

