/*
@author 如梦一般
@date 2019-07-10 12:38
*/
package controller

import (
	common "../../common/model"
	"../../db"
	"../model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UpController struct {
	RouterGroup *gin.RouterGroup
}

/*
分页的方式列举出特定状态的up主
*/
func (up *UpController) UpList() {
	up.RouterGroup.GET("/up/page/:page/status/:status", func(context *gin.Context) {
		page, _ := strconv.Atoi(context.Param("page"))
		status := context.Param("status")
		//SELECT * FROM bbd_up WHERE   bbd_up.status =0 LIMIT 20 OFFSET 2
		//sql := `select * from bbd.bbd_up where status =` + status + " limit 20 offset " + strconv.Itoa(page*20)
		psql := `select * from bbd.bbd_up where status =` + "? limit 20 offset ?"

		result := common.ApiModel{}

		if db.CheckDB() {
			stmt, err := db.FetchDB().Prepare(psql)
			ups := make([]model.Up, 0)

			if err != nil {
				fmt.Println(err.Error())
				result.Msg = err.Error()
				result.Code = common.ErrorCode
			} else {
				rows, err := stmt.Query(status, strconv.Itoa(page*20))
				if err != nil {
					fmt.Println(err.Error())
					result.Code = common.ErrorCode
					result.Msg = err.Error()
				} else {

					for rows.Next() {
						up := model.Up{}
						rows.Scan(&up.Id, &up.Mid, &up.Status, &up.Face, &up.Name)
						ups = append(ups, up)
					}
					result.Result = ups
				}
			}
		}
		context.JSONP(http.StatusOK, result)
	})
}

/*
修改up主的状态
不检测是否存在up
*/
func (up *UpController) ChangeUpStaus() {
	up.RouterGroup.GET("/up/change/:mid/status/:status", func(context *gin.Context) {
		mid := context.Param("mid")
		status := context.Param("status")
		//找到up
		//sql := `update bbd.bbd_up  bp  set bp.status=:status where bp.id=? AND bp.status =0`
		sql := `update bbd.bbd_up  bp  set bp.status=? where bp.mid=? AND bp.status =0`
		pre, err := db.FetchDB().Prepare(sql)
		result := common.ApiModel{}
		if err != nil {
			result.Code = common.ErrorCode
			result.Msg = err.Error()
		} else {
			defer pre.Close()
			res, err := pre.Exec(status, mid)
			if err != nil {
				result.Code = common.ErrorCode
				result.Msg = err.Error()
			} else {
				affect, err := res.RowsAffected()
				if err != nil {
					result.Code = common.ErrorCode
					result.Msg = err.Error()
				} else {

					result.Result = affect
					if affect == 0 {
						result.Code = common.RefusedCode
						result.Msg = "状态不能重复设定"
					} else {
						result.Msg = "success"

					}
				}
			}
		}
		context.JSONP(http.StatusOK, result)
	})
}
