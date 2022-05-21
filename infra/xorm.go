package infra

import (
	"log"

	"github.com/kkjoker/recipe/model"

	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
)

func DBInit() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:root@(db:3306)/recipe?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	// xormで使ったSQLをログに吐きます
	engine.ShowSQL(true)

	// ユーザーテーブルが存在するかチェック
	exist, err := engine.IsTableExist("users")
	if err != nil {
		log.Fatal(err)
	}

	// 存在しなければテーブルを作成します
	if !exist {
		engine.CreateTables(&model.Users{})
	}
	
	return engine
}
