/*
@author 如梦一般
@date 2019-07-10 15:08
*/
package web

import (
	"../common/model"
	"../common/model/web"
	"../db"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Run(engine *gin.Engine) {
	//设置模版位置
	engine.LoadHTMLGlob("./web/views/**/*")

	//配置web服务 添加middleware
	group := engine.Group("/web")
	add(group, nil)

	group.GET("/admin/index/:aid/up/:mid", func(c *gin.Context) {
		aid := c.Param("aid")
		page := c.Param("page")
		//fmt.Println(aid, page)
		//c.Header("Access-Control-Allow-Origin", "*")
		//c.Header("Access-Control-Expose-Headers", "Content-Length,Content-Range")
		////c.Header("Referer","https://www.bilibili.com/video/av53178281/?p=2")
		//c.Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")
		//c.Set("sid", "damu83pv")
		//c.SetCookie("buvid3", "3F56ECAE-0C45-4A87-80C0-D6EDD5D5AD4592863infoc", 60*60*24*365*2, "/", ".bilibili.com", true, true)
		//c.SetCookie("sid", "damu83pv", 60*60*24*365*2, "/", ".bilibili.com", true, true)
		////c.SetCookie("Referer", "https://www.bilibili.com/video/av53178281/?p=2", 60*60*24*365*2, "/", ".bilibili.com", true, true)
		////c.SetCookie("CURRENT_FNVAL", string(16), 60*60*24*365*2, "/", ".bilibili.com", true, true)
		//c.SetCookie("damu83pv", "1", 60*60*24*365*2, "/", ".bilibili.com", true, true)
		type Av struct {
			Aid  string
			Page string
		}
		c.HTML(http.StatusOK, "index.html", Av{Aid: aid, Page: page})
	})
	group.GET("/users/index", func(c *gin.Context) {
		c.Header("P3P", "CP='IDC DSP COR CURa ADMa  OUR IND PHY ONL COM STA'")
		c.Header("Access-Control-Allow-Origin", "'http://www.ibilibili.com")
		c.Header("Access-Control-Expose-Headers", "Content-Length,Content-Range")
		//c.Header("Referer","https://www.bilibili.com/video/av53178281/?p=2")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Users",
		})
	})
	group.GET("/admin/up/page/:page", func(context *gin.Context) {
		page := context.Param("page")
		p, _ := strconv.Atoi(page)
		offset := int((p - 1) * 20)
		result := model.ApiModel{}
		if db.CheckDB() {
			stmt, err := db.FetchDB().Prepare("SELECT bp.* FROM bbd_up bp  ORDER BY bp.mid limit 20 offset ?")
			if err != nil {
				result.Msg = err.Error()
				result.Code = model.ErrorCode
			} else {
				defer stmt.Close()
				rows, err := stmt.Query(offset)
				if err != nil {
					result.Code = model.ErrorCode
					result.Msg = err.Error()
				} else {
					ups := []web.Up{}

					for rows.Next() {

						up := web.Up{}
						rows.Scan(&up.Id, &up.Mid, &up.Status, &up.Face, &up.Name)

						if len(up.Name) > 0 {

							ups = append(ups, up)
						}
					}
					result.Result = ups
				}
			}
		}

		fmt.Println(result)
		context.HTML(http.StatusOK, "up.html", result)
	})
	group.GET("/admin/topic/:mid", func(context *gin.Context) {
		mid := context.Param("mid")
		result := model.ApiModel{}
		if db.CheckDB() {
			stmt, err := db.FetchDB().Prepare("SELECT bt.mid,bt.aid,bt.title,bt.pic,bt.description,bt.status FROM bbd_topic bt WHERE bt.mid = ?")
			if err != nil {
				result.Msg = err.Error()
				result.Code = model.ErrorCode
			} else {
				defer stmt.Close()
				rows, err := stmt.Query(mid)
				if err != nil {
					result.Code = model.ErrorCode
					result.Msg = err.Error()
				} else {
					topics := []web.Topic{}

					for rows.Next() {

						topic := web.Topic{}
						rows.Scan(&topic.Mid, &topic.Aid, &topic.Title, &topic.Pic, &topic.Description, &topic.Status)

						if len(topic.Title) > 0 {

							topics = append(topics, topic)
						}
					}
					result.Result = topics
				}
			}
		}

		fmt.Println(result)
		context.HTML(http.StatusOK, "topic.html", result)
	})
}
func add(group *gin.RouterGroup, middlewares ...gin.HandlerFunc) {
	for _, middleware := range middlewares {
		if middleware == nil {
			continue
		}
		group.Use(middleware)

	}
}
