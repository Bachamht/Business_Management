package main

import (
	"Business_Management/model"
	"Business_Management/routes"
)

func main() {
	//连接数据库
	model.ConnectDB()
	//初始化路由
	routes.InitRouter()
}
