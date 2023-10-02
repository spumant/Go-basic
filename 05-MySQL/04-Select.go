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
//// 单行查询
//func queryRowDemo() {
//	sqlStr := "select id,username,password from user_tbl where id=?"
//	var u user
//
//	// 确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
//	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.username, &u.password)
//	if err != nil {
//		fmt.Println("scan failed", err)
//		return
//	}
//	fmt.Println("id:", u.id, "name:", u.username, "password:", u.password)
//}
//
//func queryMany() {
//	sqlStr := "select id,username,password from user_tbl where id > ?"
//	query, err := db.Query(sqlStr, 0)
//	if err != nil {
//		fmt.Println("query failed :", err)
//		return
//	}
//	// 关闭query释放持有的数据库连接
//	defer query.Close()
//
//	for query.Next() {
//		var u user
//		err := query.Scan(&u.id, &u.username, &u.password)
//		if err != nil {
//			fmt.Println("scan failed :", err)
//			return
//		}
//		fmt.Println("id:", u.id, "name:", u.username, "password:", u.password)
//	}
//}
//
//func main() {
//	err := initDB()
//	if err != nil {
//		fmt.Println("初始化失败", err)
//	} else {
//		fmt.Println("初始化成功")
//	}
//	queryRowDemo()
//	queryMany()
//}
