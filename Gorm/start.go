package main

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type UserInfo struct {
	Id     uint
	Name   string
	Gender string
	Hobby  string
}

// 声明模型
type User struct {
	ID           uint
	Name         string
	Email        string
	Age          uint8
	Birthday     time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type User2 struct {
	gorm.Model
	Name string
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(&UserInfo{})
	//res := db.Create(UserInfo{
	//	Id:     2,
	//	Name:   "张大山",
	//	Gender: "male",
	//	Hobby:  "sing",
	//})
	//println(res.RowsAffected)

	//case2: query
	res2 := db.Find(&UserInfo{})
	println(res2.RowsAffected)

	// not work todo
	//user := User{Name: "张三", Age: 18, Birthday: time.Now()}
	//result := db.Create(&user) // 通过数据的指针来创建

}
