package service

import (
	"go_study_im/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

var DbEngine *xorm.Engine

func init()  {
	var err error
	DbEngine,err = xorm.NewEngine("mysql","root:123456@(192.168.0.61:3306)/go_im?charset=utf8mb4")
	if err != nil{
		log.Fatal(err.Error())
	}
	//显示打印的SQL
	DbEngine.ShowSQL(true)
	DbEngine.SetMaxOpenConns(2)
	fmt.Println("xorm init success")
	DbEngine.Sync2(model.User{},model.Community{},model.Contact{})
}

