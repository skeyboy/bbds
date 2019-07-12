/*
@author 如梦一般
@date 2019-07-10 11:20
*/
package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"time"
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
	db, err := sql.Open("mysql", "root:@/bbd")
	if err != nil {
		_HasDB = false
	} else {
		_HasDB = true
	}
	_DB = db
	_DB.SetConnMaxLifetime(time.Second * 30)
	_DB.SetMaxIdleConns(30)
	_DB.SetMaxOpenConns(100)
}
func CheckDB() bool {
	return _HasDB
}
func FetchDB() *sql.DB {
	return _DB
}
func Close() {
	lock.Lock()
	defer lock.Unlock()
	if CheckDB() {
		_DB.Close()
	}
}
