package sqlite

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var x *xorm.Engine

func init() {
	var err error
	//创建ORM引擎与数据库
	x, err = xorm.NewEngine("sqlite3", "./sql/xhcms.db")
	if err != nil {
		log.Fatalf("Fail to create engine: %v\n", err)
	}
	//同步结构体与数据表
	if err = x.Sync2(new(User)); err != nil {
		log.Fatalf("Fail to sync database: %v\n", err)
	}

}
