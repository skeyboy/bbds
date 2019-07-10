/*
@author 如梦一般
@date 2019-07-10 15:11
*/
package start

import (
	API "../api"
	Web "../web"
)
import "github.com/gin-gonic/gin"

func Run(run func(engin *gin.Engine), middlewares []gin.HandlerFunc) {
	engin := gin.Default()
	addMiddleware(engin, middlewares)
	API.Run(engin)
	Web.Run(engin)
	run(engin)
}

//使用全局的middleware
func addMiddleware(engin *gin.Engine, middlewares []gin.HandlerFunc) {
	for _, middleware := range middlewares {
		if middleware == nil {
			continue
		}
		engin.Use(middleware)
	}
}
