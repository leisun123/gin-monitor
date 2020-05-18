package routes

import (
	"gin-monitor/model"
	"gin-monitor/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Get 获取任务列表
// @Tags Task
// @Summary 获取任务列表
// @Success 200 {object} string
// @Failure 401 {object} model.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 404 {object} model.HTTPError "{error:{code:0,message:资源不存在}}"
// @Failure 500 {object} model.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router /task [get]
func GetTaskList(c *gin.Context) {

	results, err := model.GetTaskList()
	if err != nil {
		HandleError(http.StatusInternalServerError, c, err)
		return
	}
	HandleSuccessData(c, results)
}


// PUT 创建新任务
// @Tags Task
// @Summary 创建新任务
// @Param TaskDto body model.TaskDto true "新任务详情"
// @Success 200 {object} string
// @Failure 401 {object} model.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 404 {object} model.HTTPError "{error:{code:0,message:资源不存在}}"
// @Failure 500 {object} model.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router /task [put]
func NewTask(c *gin.Context)  {
	var item model.TaskDto
	// 绑定数据模型
	if err := c.ShouldBindJSON(&item); err != nil {
		HandleError(http.StatusBadRequest, c, err)
		return
	}

	// 更新数据库
	if err := model.AddTask(item); err != nil {
		HandleError(http.StatusInternalServerError, c, err)
		return
	}
}

// DELETE 删除任务
// @Tags Task
// @Summary 删除任务
// @Param task_id path string true "任务ID"
// @Success 200 {object} string
// @Failure 401 {object} model.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 404 {object} model.HTTPError "{error:{code:0,message:资源不存在}}"
// @Failure 500 {object} model.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router /task/{task_id} [delete]
func RemoveTaskById(c *gin.Context)  {
	taskId := c.Param("task_id")
	// 删除定时任务
	if err := model.RemoveTask(taskId); err != nil {
		HandleError(http.StatusInternalServerError, c, err)
		return
	}

	// 更新定时任务
	if err := services.Sched.Update(); err != nil {
		HandleError(http.StatusInternalServerError, c, err)
		return
	}
}

// POST 修改任务
// @Tags Task
// @Summary 修改任务
// @Param task_id path string true "任务ID"
// @Param task body model.TaskDto true "任务详情"
// @Success 200 {object} string
// @Failure 401 {object} model.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 404 {object} model.HTTPError "{error:{code:0,message:资源不存在}}"
// @Failure 500 {object} model.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router /task/{task_id} [post]
func UpdateTask(c *gin.Context)  {
	taskId := c.Param("task_id")
	var item model.TaskDto
	// 绑定数据模型
	if err := c.ShouldBindJSON(&item); err != nil {
		HandleError(http.StatusBadRequest, c, err)
		return
	}

	// 更新数据库
	if err := model.UpdateTask(taskId, item); err != nil {
		HandleError(http.StatusInternalServerError, c, err)
		return
	}

	// 更新定时任务
	if err := services.Sched.Update(); err != nil {
		HandleError(http.StatusInternalServerError, c, err)
		return
	}
}


// Get 获取定时任务列表
// @Tags Schedule
// @Summary 获取定时任务列表
// @Success 200 {object} string
// @Failure 401 {object} model.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 404 {object} model.HTTPError "{error:{code:0,message:资源不存在}}"
// @Failure 500 {object} model.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router /schedules [get]
func GetScheduleList(c *gin.Context) {

	results, err := model.GetScheduleList()
	if err != nil {
		HandleError(http.StatusInternalServerError, c, err)
		return
	}
	HandleSuccessData(c, results)
}


// Get 获取定时任务详情
// @Tags Schedule
// @Summary 获取定时任务详情
// @Param schedule_id path string true "定时任务ID"
// @Success 200 {object} string
// @Failure 401 {object} model.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 404 {object} model.HTTPError "{error:{code:0,message:资源不存在}}"
// @Failure 500 {object} model.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router /schedules/{schedule_id} [get]
func GetSchedule(c *gin.Context) {
	id := c.Param("id")

	result, err := model.GetSchedule(id)
	if err != nil {
		HandleError(http.StatusInternalServerError, c, err)
		return
	}

	HandleSuccessData(c, result)
}


// POST 修改定时任务
// @Tags Schedule
// @Summary 修改定时任务
// @Param schedule_id path string true "定时任务ID"
// @Param ScheduleDto body model.ScheduleDto true "定时任务详情"
// @Success 200 {object} string
// @Failure 401 {object} model.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 404 {object} model.HTTPError "{error:{code:0,message:资源不存在}}"
// @Failure 500 {object} model.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router /schedules/{schedule_id} [post]
func PostSchedule(c *gin.Context) {
	id := c.Param("id")

	// 绑定数据模型
	var newItem model.ScheduleDto
	if err := c.ShouldBindJSON(&newItem); err != nil {
		HandleError(http.StatusBadRequest, c, err)
		return
	}

	// 验证cron表达式
	if err := services.ParserCron(newItem.Cron); err != nil {
		HandleError(http.StatusInternalServerError, c, err)
		return
	}

	// 更新数据库
	if err := model.UpdateSchedule(id, newItem); err != nil {
		HandleError(http.StatusInternalServerError, c, err)
		return
	}

	// 更新定时任务
	if err := services.Sched.Update(); err != nil {
		HandleError(http.StatusInternalServerError, c, err)
		return
	}

	HandleSuccess(c)
}


// PUT 新建定时任务
// @Tags Schedule
// @Summary 新建定时任务
// @Param ScheduleDto body model.ScheduleDto true "定时任务信息"
// @Success 200 {object} string
// @Failure 401 {object} model.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 404 {object} model.HTTPError "{error:{code:0,message:资源不存在}}"
// @Failure 500 {object} model.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router /schedules [put]
func PutSchedule(c *gin.Context) {
	var item model.ScheduleDto

	// 绑定数据模型
	if err := c.ShouldBindJSON(&item); err != nil {
		HandleError(http.StatusBadRequest, c, err)
		return
	}

	// 验证cron表达式
	if err := services.ParserCron(item.Cron); err != nil {
		HandleError(http.StatusInternalServerError, c, err)
		return
	}


	// 更新数据库
	if err := model.AddSchedule(item); err != nil {
		HandleError(http.StatusInternalServerError, c, err)
		return
	}

	// 更新定时任务
	if err := services.Sched.Update(); err != nil {
		HandleError(http.StatusInternalServerError, c, err)
		return
	}

	HandleSuccess(c)
}


// Delete 删除定时任务
// @Tags Schedule
// @Summary 删除定时任务
// @Param schedule_id path string true "定时任务ID"
// @Success 200 {object} string
// @Failure 401 {object} model.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 404 {object} model.HTTPError "{error:{code:0,message:资源不存在}}"
// @Failure 500 {object} model.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router /schedules/{schedule_id} [delete]
func DeleteSchedule(c *gin.Context) {
	id := c.Param("id")

	// 删除定时任务
	if err := model.RemoveSchedule(id); err != nil {
		HandleError(http.StatusInternalServerError, c, err)
		return
	}

	// 更新定时任务
	if err := services.Sched.Update(); err != nil {
		HandleError(http.StatusInternalServerError, c, err)
		return
	}

	HandleSuccess(c)
}


// POST 停止定时任务
// @Tags Schedule
// @Summary 停止定时任务
// @Param schedule_id path string true "定时任务ID"
// @Success 200 {object} string
// @Failure 401 {object} model.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 404 {object} model.HTTPError "{error:{code:0,message:资源不存在}}"
// @Failure 500 {object} model.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router /schedules/{schedule_id}/disable [post]
func DisableSchedule(c *gin.Context) {
	id := c.Param("id")
	if err := services.Sched.Disable(id); err != nil {
		HandleError(http.StatusInternalServerError, c, err)
		return
	}
	HandleSuccess(c)
}


// POST 运行定时任务
// @Tags Schedule
// @Summary 运行定时任务
// @Param schedule_id path string true "定时任务ID"
// @Success 200 {object} string
// @Failure 401 {object} model.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 404 {object} model.HTTPError "{error:{code:0,message:资源不存在}}"
// @Failure 500 {object} model.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router /schedules/{schedule_id}/enable [post]
func EnableSchedule(c *gin.Context) {
	id := c.Param("id")
	if err := services.Sched.Enable(id); err != nil {
		HandleError(http.StatusInternalServerError, c, err)
		return
	}
	HandleSuccess(c)
}

