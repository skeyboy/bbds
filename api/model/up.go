/*
@author 如梦一般
@date 2019-07-10 12:12
*/
package model

type Up struct {
	Id     int
	Mid    int
	Status int
	Face   string
	Name   string
}

func (up *Up) IsValid() bool {
	return up.Status == 0
}
