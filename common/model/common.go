/*
@author 如梦一般
@date 2019-07-10 12:13
*/
package model

import (
	"fmt"
	"strconv"
)

const (
	//默认为操作成功
	SuccessDefaultCode = 0
	//有异常
	ErrorCode = 1
	//访问拒绝、操作拒绝
	RefusedCode = 2
)

type ApiModel struct {
	Code   int
	Result interface{}
	Msg    string
}

func (m *ApiModel) String() string {

	code := strconv.Itoa(m.Code)
	return fmt.Sprintln("code:%s\tresult:%s\t\n%s\n", code, m.Result, m.Code)
}

func (m *ApiModel) Value() interface{} {
	return m.Result
}
