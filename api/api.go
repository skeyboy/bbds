/*
@author 如梦一般
@date 2019-07-10 12:31
*/
package api

import (
	"./controller"
	m "./middleware"
	"github.com/gin-gonic/gin"
)

func Run(router *gin.Engine) {

	//分组
	group := router.Group("/api")
	//中间件
	m.UsemiddlewareFor(group)

	upCtrl := controller.UpController{RouterGroup: group}
	upCtrl.UpList()
	upCtrl.ChangeUpStaus()
}
