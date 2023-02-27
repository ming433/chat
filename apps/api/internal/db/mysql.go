package db

import (
	"chat/apps/api/global"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitMysql() {
	dbConfig := global.GlobalConfig
	gormConfig := &gorm.Config{
		SkipDefaultTransaction: dbConfig.Gorm.SkipDefaultTx,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   dbConfig.Gorm.TablePrefix,
			SingularTable: dbConfig.Gorm.SingularTable,
		},
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database)
	conn, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		fmt.Println("create mysql client failed.")
		return
	}
	mysqlDB, err := conn.DB()
	if err != nil {
		fmt.Println("Invalid DB.")
		return
	}
	mysqlDB.SetMaxIdleConns(10)
	mysqlDB.SetMaxOpenConns(100)
	if err = mysqlDB.Ping(); err != nil {
		fmt.Println("mysql client failed.")
	}
	fmt.Println("mysql client success.")
	global.DB = conn
	fmt.Println("mysql client success.", global.DB)
}
