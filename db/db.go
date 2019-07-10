/*
@author 如梦一般
@date 2019-07-10 11:20
*/
package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

var _DB *sql.DB
var _HasDB bool = false

var once sync.Once = sync.Once{}
var lock sync.RWMutex

func init() {
	once.Do(bbdDB)
}
func bbdDB() {
	lock.Lock()
	defer lock.Unlock()
	db, err := sql.Open("mysql", "root:12345678@/bbd")
	if err != nil {
		_HasDB = false
	} else {
		_HasDB = true
	}
	_DB = db
}
func CheckDB() bool {
	return _HasDB
}
func FetchDB() *sql.DB {
	return _DB
}
