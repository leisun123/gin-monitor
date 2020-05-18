/*
Package routers 生成swagger文档

文档规则请参考：https://github.com/swaggo/swag#declarative-comments-format

使用方式：

	go get -u github.com/swaggo/swag/cmd/swag
	swag init -g ./routes/swagger.go -o ./main/docs/swagger*/
package routes

// @title op-robot
// @version 5.2.1
// @description RBAC scaffolding based on Gin + Gorm + Casbin + Dig.
// @schemes http https
// @host
// @basePath /