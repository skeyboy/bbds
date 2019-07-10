package main

import (
	Start "./start"
	"github.com/gin-gonic/gin"
)

func start() {
	Start.Run(run, middlewares())
}

//启动服务
func run(engin *gin.Engine) {
	globalConfig(engin)
	err := engin.Run(":8080")
	if err != nil {
		panic(err)
	}
}

//全局中间件
func middlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{}
}

//此处进行一些全局性设置
func globalConfig(engin *gin.Engine) {

}

//启动入口
func main() {
	start()
}
