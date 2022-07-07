package Databases

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func InitDB() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=true&loc=Local",
		"root", "root", "127.0.0.1:3306", "test", "utf8")

	if conn, err := sql.Open("mysql", dsn); err != nil {
		panic(err.Error())
	} else {
		conn.SetConnMaxLifetime(7 * time.Second) //设置空闲时间，这个是比mysql 主动断开的时候短
		conn.SetMaxOpenConns(10)
		conn.SetMaxIdleConns(10)
		return conn
	}
}