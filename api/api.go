/*
@author 如梦一般
@date 2019-07-10 12:31
*/
package api

import (
	"./controller"
	"github.com/gin-gonic/gin"
)

func RunApi(router *gin.Engine) {

	upCtrl := controller.UpController{Router: router}
	upCtrl.UpList()
	upCtrl.ChangeUpStaus()
}
