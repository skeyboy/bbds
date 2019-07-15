/*
@author 如梦一般
@date 2019-07-10 14:50
*/
package middleware

import "github.com/gin-gonic/gin"

func UsemiddlewareFor(group *gin.RouterGroup) {
	group.Use(func(c *gin.Context) {
		c.Header("P3P", "CP='IDC DSP COR CURa ADMa  OUR IND PHY ONL COM STA'")
		c.Header("Access-Control-Allow-Origin", "https://www.bilibili.com/")
		c.Header("referrer", "https://www.bilibili.com/")
	})
}
