package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	//router.LoadHTMLGlob("./views/**/*")
	router.LoadHTMLFiles("./views/index.html")
	router.GET("/index/:aid/page/:page", func(c *gin.Context) {
		aid := c.Param("aid")
		page := c.Param("page")
		fmt.Println(aid, page)
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Expose-Headers", "Content-Length,Content-Range")
		//c.Header("Referer","https://www.bilibili.com/video/av53178281/?p=2")
		c.Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")
		c.Set("sid", "damu83pv")
		c.SetCookie("buvid3", "3F56ECAE-0C45-4A87-80C0-D6EDD5D5AD4592863infoc", 60*60*24*365*2, "/", ".bilibili.com", true, true)
		c.SetCookie("sid", "damu83pv", 60*60*24*365*2, "/", ".bilibili.com", true, true)
		//c.SetCookie("Referer", "https://www.bilibili.com/video/av53178281/?p=2", 60*60*24*365*2, "/", ".bilibili.com", true, true)
		//c.SetCookie("CURRENT_FNVAL", string(16), 60*60*24*365*2, "/", ".bilibili.com", true, true)
		c.SetCookie("damu83pv", "1", 60*60*24*365*2, "/", ".bilibili.com", true, true)
		type Av struct {
			Aid  string
			Page string
		}
		c.HTML(http.StatusOK, "index.html", Av{Aid: aid, Page: page})
	})
	router.GET("/users/index", func(c *gin.Context) {
		c.Header("P3P", "CP='IDC DSP COR CURa ADMa  OUR IND PHY ONL COM STA'")
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})
	router.Run(":8080")
}
