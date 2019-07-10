/*
@author 如梦一般
@date 2019-07-10 14:50
*/
package middleware

import "github.com/gin-gonic/gin"

func UsemiddlewareFor(group *gin.RouterGroup) {
	group.Use(func(context *gin.Context) {

	})
}
