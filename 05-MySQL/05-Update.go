package main

//import (
//	"database/sql"
//	"fmt"
//	_ "github.com/go-sql-driver/mysql"
//)
//
//var db *sql.DB
//
//type user struct {
//	id       int
//	username string
//	password string
//}
//
//func initDB() (err error) {
//	dsn := "root:monarch99@tcp(127.0.0.1:3306)/go_db?charset=utf8&parseTime=True"
//	db, err = sql.Open("mysql", dsn)
//	if err != nil {
//		return err
//	}
//	// 尝试与数据库建立连接
//	err = db.Ping()
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func updateData() {
//	sqlStr := "update user_tbl set username=?,password=? where id=?"
//	ret, err := db.Exec(sqlStr, "张三2", "zs2123", 1)
//	if err != nil {
//		fmt.Println("update failed:", err)
//		return
//	}
//	row, err := ret.RowsAffected()
//	if err != nil {
//		fmt.Println("update row failed ：", err)
//		return
//	}
//	fmt.Println("success", row)
//}
//
//func main() {
//	err := initDB()
//	if err != nil {
//		fmt.Println("初始化失败", err)
//	} else {
//		fmt.Println("初始化成功")
//	}
//	updateData()
//}
