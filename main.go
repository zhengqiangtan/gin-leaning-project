package main

import "gin-leaning-project/Databases"

/**
  case1: https://www.jianshu.com/p/92919004293d
  case2: https://cloud.tencent.com/developer/article/1596713
  case3: mysql crud https://www.yangyanxing.com/article/use-go-gin-mysql-base.html  赞
 */
func main() {
	db := Databases.InitDB()

	//rows, err :=db.Query("select * from user")
	//for rows.Next() {
	//	err := rows.Scan(&ID, &a.Name, &a.Age)
	//	if err != nil {
	//		print(rows.Next())
	//	}
	//}

	println(db.Ping())
	print(db)

}
