package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	Id   uint64 `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Age  uint8  `json:"age" db:"age"`
}

var db *sqlx.DB

// initializeDatabase 初始化连接数据库
func initializeDatabase() (err error) {
	dsn := "root:Hunan19970708!@/student?charset=utf8mb4&parseTime=True&loc=Local"
	// Connect to a database and verify with a ping.
	// Open ---> Ping
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect to database failed, err: %v \n", err)
		return err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return nil
}

// querySingleRow
func querySingleRow() {
	sql := "select * from student where id = ?"

	var u User

	// 这里的 Get 已经内置了 scan，不需要再自己去调用 scan
	if err := db.Get(&u, sql, 1); err != nil {
		fmt.Printf("query failed err: %v \n", err)
		return
	}

	fmt.Printf("id: %d, name: %s, age: %d", u.Id, u.Name, u.Age)
}

// queryMultiRow
func queryMultiRow() []User {
	sql := "select * from student"
	var users []User

	if err := db.Select(&users, sql); err != nil {
		fmt.Printf("query failed err: %v \n", err)
		return nil
	}

	return users
}

// updateRow
func updateRow() {
	sql := "update student set age = ? where id = ?"
	res, err := db.Exec(sql, 10, 2)
	if err != nil {
		fmt.Printf("update failed, err = %v \n", err)
		return
	}

	n, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("update failed, err = %v \n", err)
		return
	}

	fmt.Printf("update success, n = %v \n", n)
}

func insertRow() {
	sql := "insert into student (name, age) values(?, ?)"

	res, err := db.Exec(sql, "孙", 20)
	if err != nil {
		fmt.Printf("插入数据失败：%v \n", err)
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("get LastInsertId failed: %v \n", id)
	}

	fmt.Printf("插入数据成功：%v \n", id)
}

// deleteRow
func deleteRow() {
	sql := "delete from student where id = ?"
	res, err := db.Exec(sql, 10)
	if err != nil {
		fmt.Printf("delete failed err: %v\n", err)
		return
	}

	n, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("delete RowsAffected failed err: %v\n", err)
		return
	}

	fmt.Printf("delete success affected row = %d \n", n)
}
func main() {
	if err := initializeDatabase(); err != nil {
		panic(err)
	}

	fmt.Println("connect to db successful")

	//querySingleRow()
	//r := gin.Default()
	//r.GET("user", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"data": queryMultiRow(),
	//	})
	//})
	//r.Run()

	//updateRow()
	//insertRow()

	deleteRow()
}
