/**
 * @author: hqd
 * @description: mysql
 * @file: mysql
 * @date: 2021-02-05 11:54
 */

package services

import (
	"fmt"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var dbMaster *gorm.DB
var dbSlave *gorm.DB

func InitDB() error {
	mysqlUser, _ := web.AppConfig.String("mysqluser")
	mysqlPass, _ := web.AppConfig.String("mysqlpass")
	mysqlHost, _ := web.AppConfig.String("mysqlhost")
	mysqlPort, _ := web.AppConfig.String("mysqlport")
	mysqlDB, _ := web.AppConfig.String("mysqldb")
	mysqlMaxOpenConn, _ := web.AppConfig.Int("maxopenconn")
	mysqlMaxIdleConn, _ := web.AppConfig.Int("maxidleconn")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		mysqlUser,
		mysqlPass,
		mysqlHost,
		mysqlPort,
		mysqlDB,
	)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		logs.Error("connection to db err:%v", err)
		return err
	}

	if web.BConfig.RunMode == web.DEV {
		db.LogMode(true)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "t_" + defaultTableName
	}
	db.SingularTable(true)
	db.DB().SetMaxOpenConns(mysqlMaxOpenConn)
	db.DB().SetMaxIdleConns(mysqlMaxIdleConn)
	setDBConnection(db, db)
	return nil
}

func setDBConnection(master, slave *gorm.DB) {
	if slave == nil {
		slave = master
	}
	dbMaster = master
	dbSlave = slave
	logs.Info("init db connection success")
}
