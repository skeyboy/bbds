/*
@author 如梦一般
@date 2019-07-10 15:08
*/
package web

import (
	"../common/model"
	"../common/model/web"
	"../db"
	m "./middleware"
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"

	"net/http"
	"path/filepath"
	"strconv"
)

const (
	TemplateJoin = "@"
)

func loadTemplates(templatesDir string, categories []string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	for _, category := range categories {
		layouts, err := filepath.Glob(templatesDir + "/layouts/" + category + "-*.html")
		if err != nil {
			panic(err.Error())
		}

		includes, err := filepath.Glob(templatesDir + "/includes/" + category + "/*.html")
		if err != nil {
			panic(err.Error())
		}

		// Generate our templates map from our layouts/ and includes/ directories
		for _, include := range includes {
			layoutCopy := make([]string, len(layouts))
			copy(layoutCopy, layouts)
			files := append(layoutCopy, include)
			r.AddFromFiles(category+TemplateJoin+filepath.Base(include), files...)
		}
	}

	return r
}
func Run(engine *gin.Engine) {
	//设置模版位置
	//engine.LoadHTMLGlob("./web/views/**/*")
	group := engine.Group("/web")
	engine.HTMLRender = loadTemplates("./web/views", []string{"admin", "member"})
	//engine.Static("/resources", "./web")

	//配置web服务 添加middleware
	m.UsemiddlewareFor(group)
	add(group, nil)
	group.GET("/category", func(context *gin.Context) {
		result := model.ApiModel{}

		rows, err := db.FetchDB().Query(`select c.id, c.category_name name from categories c`)
		if err != nil {

		} else {
			type Category struct {
				Name string
				Id   string
			}
			var cats = []Category{}

			for rows.Next() {
				var cat = Category{}

				rows.Scan(&cat.Id, &cat.Name)

				cats = append(cats, cat)
			}
			if len(cats) == 0 {

				cats = append(cats, Category{Name: "其他", Id: "0"})
				cats = append(cats, Category{Name: "其他", Id: "0"})
				cats = append(cats, Category{Name: "其他", Id: "0"})
			}
			result.Result = cats
		}
		context.HTML(http.StatusOK, "admin"+TemplateJoin+"category.html", result)
	})
	group.GET("/admin/index/av/:aid/up/:mid", func(c *gin.Context) {
		type Album struct {
			Videos int
			Title  string
			Origin string
			Aid    int64
		}
		type Av struct {
			Aid   string
			Page  string
			Album Album
		}
		aid := c.Param("aid")
		sql := `select ba.videos, ba.title, ba.origin,ba.aid from bbd_album ba where ba.aid =?`
		av := Av{Aid: aid, Page: "1"}

		if db.CheckDB() {
			stmt, _ := db.FetchDB().Prepare(sql)
			row := stmt.QueryRow(aid)
			row.Scan(&av.Album.Videos, &av.Album.Title, &av.Album.Origin, &av.Album.Aid)
		}
		result := model.ApiModel{}
		result.Result = av

		c.HTML(http.StatusOK, "admin"+TemplateJoin+"index.html", result)
	})
	group.GET("/users/index", func(c *gin.Context) {
		c.Header("P3P", "CP='IDC DSP COR CURa ADMa  OUR IND PHY ONL COM STA'")
		c.Header("Access-Control-Allow-Origin", "'http://www.ibilibili.com")
		c.Header("Access-Control-Expose-Headers", "Content-Length,Content-Range")
		//c.Header("Referer","https://www.bilibili.com/video/av53178281/?p=2")
		c.HTML(http.StatusOK, "member"+TemplateJoin+"index.html", gin.H{
			"title": "Users",
		})
	})
	group.GET("/admin/up/page/:page", func(context *gin.Context) {
		page := context.Param("page")
		p, _ := strconv.Atoi(page)
		offset := int((p - 1) * 20)
		type Result struct {
			model.ApiModel
			Page int
		}

		result := Result{}
		result.Page = p
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
						rows.Scan(&up.Id, &up.Mid, &up.Status, &up.Face, &up.Name, &up.Date)

						if len(up.Name) > 0 {

							ups = append(ups, up)
						}
					}
					result.Result = ups
				}
			}
		}

		fmt.Println(result)
		context.HTML(http.StatusOK, "admin"+TemplateJoin+"up.html", result)
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
		context.HTML(http.StatusOK, "admin"+TemplateJoin+"topic.html", result)
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
