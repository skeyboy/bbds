package main

import (
	"./db"
	Start "./start"
	"github.com/gin-gonic/gin"
)

func start() {
	Start.Run(run, middlewares())
	defer func() {
		if db.CheckDB() {
			db.Close()
		}
	}()
}

//启动服务
func run(engin *gin.Engine) {
	globalConfig(engin)
	err := engin.Run(":8081")
	if err != nil {
		panic(err.Error())
	}
}

//全局中间件
func middlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{}
}

//此处进行一些全局性设置
func globalConfig(engin *gin.Engine) {
	//engin.Static("/resources","./web")
	//engin.StaticFS("/resources",http.Dir("./web"))
}

//启动入口
func main() {
	start()
}
