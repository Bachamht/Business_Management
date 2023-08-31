package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"strings"
)

/*
完成数据库启动以及一些数据库相关的操作
*/
var DB *sql.DB
var err error

func ConnectDB() {

	Mysqlinfo := GetDbInfo()
	dbURL := Mysqlinfo.Mysql.User + ":" + Mysqlinfo.Mysql.Password + "@tcp(" + Mysqlinfo.Mysql.Host + ":" + Mysqlinfo.Mysql.Port + ")/" + Mysqlinfo.Mysql.Dbname

	DB, err = sql.Open("mysql", dbURL)
	if err != nil {
		fmt.Println(err)
	}

	//设置上数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)

	if err := DB.Ping(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("connnect success")

	// 读取 SQL 脚本内容
	sqlScript, err := ioutil.ReadFile("database/db.sql")
	if err != nil {
		log.Fatal(err)
	}

	// 执行 SQL 脚本

	statements := strings.Split(string(sqlScript), ";")

	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt != "" {
			_, err := DB.Exec(stmt)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	fmt.Println("SQL script executed successfully.")

}
