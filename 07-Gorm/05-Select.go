package main

//import (
//	"fmt"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//	"gorm.io/gorm/clause"
//	"log"
//	"time"
//)
//
//var db *gorm.DB
//
//func init() {
//	dsn := "root:monarch99@tcp(127.0.0.1:3306)/go_db?charset=utf8&parseTime=True"
//	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		log.Fatal(err)
//	}
//	db = d
//}
//
//type User struct {
//	gorm.Model
//	Name     string
//	Age      int
//	Birthday time.Time
//}
//
//func createTable() {
//	db.AutoMigrate(&User{})
//}
//
//func select1() {
//	var user User
//	db.First(&user) // 按主键排序第一个 未指定主键则按照第一个字段排序
//	fmt.Println(user)
//	db.Take(&user) // 第一个
//	db.Last(&user) // 最后一个
//}
//
//func select2() {
//	var user User
//	db.First(&user, 2)
//	fmt.Println(user.ID)
//
//	var users []User
//	db.Find(&users, []int{1, 2, 3})
//	//fmt.Println(users)
//	for _, v := range users {
//		fmt.Println(v)
//	}
//}
//
//func select3() {
//	var users []User
//	result := db.Find(&users) // 查询所有
//	fmt.Println(result.RowsAffected)
//}
//
//func select4() {
//	var user User
//	var users []User
//
//	lastWeek := time.Now()
//	today := time.Now()
//
//	// 获取第一条匹配的记录
//	db.Where("name = ?", "jinzhu").First(&user)
//	// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;
//
//	// 获取全部匹配的记录
//	db.Where("name <> ?", "jinzhu").Find(&users)
//	// SELECT * FROM users WHERE name <> 'jinzhu';
//
//	// IN
//	db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)
//	// SELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');
//
//	// LIKE
//	db.Where("name LIKE ?", "%jin%").Find(&users)
//	// SELECT * FROM users WHERE name LIKE '%jin%';
//
//	// AND
//	db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
//	// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;
//
//	// Time
//	db.Where("updated_at > ?", lastWeek).Find(&users)
//	// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';
//
//	// BETWEEN
//	db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
//	// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';
//
//	// Struct
//	db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
//	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;
//
//	// Map
//	db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
//	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;
//
//	// 主键切片条件
//	db.Where([]int64{20, 21, 22}).Find(&users)
//	// SELECT * FROM users WHERE id IN (20, 21, 22);
//
//	db.Where(&User{Name: "jinzhu", Age: 0}).Find(&users)
//	// SELECT * FROM users WHERE name = "jinzhu";
//
//	db.Where(map[string]interface{}{"Name": "jinzhu", "Age": 0}).Find(&users)
//	// SELECT * FROM users WHERE name = "jinzhu" AND age = 0;
//
//}
//
//func select5() {
//	var user User
//	var users []User
//
//	db.Not("name = ?", "jinzhu").First(&user)
//	// SELECT * FROM users WHERE NOT name = "jinzhu" ORDER BY id LIMIT 1;
//
//	// Not In
//	db.Not(map[string]interface{}{"name": []string{"jinzhu", "jinzhu 2"}}).Find(&users)
//	// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");
//
//	// Struct
//	db.Not(User{Name: "jinzhu", Age: 18}).First(&user)
//	// SELECT * FROM users WHERE name <> "jinzhu" AND age <> 18 ORDER BY id LIMIT 1;
//
//	// 不在主键切片中的记录
//	db.Not([]int64{1, 2, 3}).First(&user)
//	// SELECT * FROM users WHERE id NOT IN (1,2,3) ORDER BY id LIMIT 1;
//}
//
//func select6() {
//	var users []User
//
//	db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)
//	// SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';
//
//	// Struct
//	db.Where("name = 'jinzhu'").Or(User{Name: "jinzhu 2", Age: 18}).Find(&users)
//	// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);
//
//	// Map
//	db.Where("name = 'jinzhu'").Or(map[string]interface{}{"name": "jinzhu 2", "age": 18}).Find(&users)
//	// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);
//}
//
//func select7() {
//	var users []User
//
//	db.Select("name", "age").Find(&users)
//	// SELECT name, age FROM users;
//
//	db.Select([]string{"name", "age"}).Find(&users)
//	// SELECT name, age FROM users;
//
//	db.Table("users").Select("COALESCE(age,?)", 42).Rows()
//	// SELECT COALESCE(age,'42') FROM users;
//}
//
//func select8() {
//	var users []User
//
//	db.Order("age desc, name").Find(&users)
//	// SELECT * FROM users ORDER BY age desc, name;
//
//	// 多个 order
//	db.Order("age desc").Order("name").Find(&users)
//	// SELECT * FROM users ORDER BY age desc, name;
//
//	db.Clauses(clause.OrderBy{
//		Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true},
//	}).Find(&User{})
//	// SELECT * FROM users ORDER BY FIELD(id,1,2,3)
//}
//
//func select9() {
//	var users, users1, users2 []User
//
//	db.Limit(3).Find(&users)
//	// SELECT * FROM users LIMIT 3;
//
//	// 通过 -1 消除 Limit 条件
//	db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
//	// SELECT * FROM users LIMIT 10; (users1)
//	// SELECT * FROM users; (users2)
//
//	db.Offset(3).Find(&users)
//	// SELECT * FROM users OFFSET 3;
//
//	db.Limit(10).Offset(5).Find(&users)
//	// SELECT * FROM users OFFSET 5 LIMIT 10;
//
//	// 通过 -1 消除 Offset 条件
//	db.Offset(10).Find(&users1).Offset(-1).Find(&users2)
//	// SELECT * FROM users OFFSET 10; (users1)
//	// SELECT * FROM users; (users2)
//}
//
//
//func main() {
//
//}
