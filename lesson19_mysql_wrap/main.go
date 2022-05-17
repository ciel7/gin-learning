package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

// initializeDatabase 初始化连接数据库
func initializeDatabase() (err error) {
	dsn := "root:Hunan19970708!@/student?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = sql.Open("mysql", dsn)

	if err != err {
		return err
	}

	if err = db.Ping(); err != nil {
		return
	}

	return nil
}

type user struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	Age  uint32 `json:"age"`
}

// querySingleRow 查询单行数据
func querySingleRow() user {
	sql := "select * from student where id = ?"
	var u user
	if err := db.QueryRow(sql, 3).Scan(&u.Id, &u.Name, &u.Age); err != nil {
		log.Printf("scan failed err: %v \n", err)
		return u
	}

	log.Println(u.Id, "-", u.Name, "-", u.Age)
	return u
}

// queryMultiRow 查询多行数据
func queryMultiRow() []user {
	sql := "select * from student"
	rows, err := db.Query(sql)
	if err != nil {
		log.Printf("db query failed err: %v \n", err)
		return nil
	}

	defer rows.Close()

	users := make([]user, 0)
	for rows.Next() {
		var u user
		err := rows.Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			log.Println(err)
			return nil
		}

		users = append(users, u)
	}
	return users
}

// updateRow 更新数据
func updateRow() {
	sql := "update student set name = ? where id = ?"
	res, err := db.Exec(sql, "zhangshan", 1)

	if err != nil {
		fmt.Printf("update failed err: %v \n", err)
	}

	n, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed err: %v \n", err)
	}

	fmt.Printf("get RowsAffected success, n = %v \n", n)
}

// deleteRow 删除数据
func deleteRow() {
	sql := "delete from student where id = ?"
	res, err := db.Exec(sql, 1)

	if err != nil {
		fmt.Printf("delete failed err: %v \n", err)
	}

	n, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed err: %v \n", err)
	}

	fmt.Printf("get RowsAffected success, n = %v \n", n)
}

// insertRow 插入数据
func insertRow() {
	sql := "insert into student(name, age) values (?, ?)"
	res, err := db.Exec(sql, "赵", 19)

	if err != nil {
		fmt.Printf("insert failed err: %v \n", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("get LastInsertId failed err: %v \n", err)
	}

	fmt.Printf("get LastInsertId success, n = %v \n", id)
}

// queryMultiRow 查询
func queryMultiRowPrepare(name string) []user {
	sql := "select * from student where name = ?"
	stmt, err := db.Prepare(sql)

	if err != nil {
		log.Printf("query failed, err: %v \n", err)
		return nil
	}

	rows, _ := stmt.Query(name)
	defer rows.Close()

	users := make([]user, 0)
	for rows.Next() {
		var u user
		err := rows.Scan(&u.Id, &u.Name, &u.Age)

		if err != nil {
			fmt.Printf("scan failed, err: %v \n", err)
			return nil
		}
		users = append(users, u)
	}
	return users
}

func main() {
	if err := initializeDatabase(); err != nil {
		panic(err)
	}

	fmt.Println("connect to db successful")
	//updateRow()
	//deleteRow()
	// insertRow()

	users := queryMultiRowPrepare("赵")
	fmt.Println("users = ", users)

	//r := gin.Default()
	//r.GET("user", func(c *gin.Context) {
	//	u := querySingleRow()
	//	c.JSON(http.StatusOK, gin.H{
	//		"u": u,
	//	})
	//})

	//r.GET("user", func(c *gin.Context) {
	//	u := queryMultiRow()
	//	c.JSON(http.StatusOK, gin.H{
	//		"u": u,
	//	})
	//})
	//r.Run()
}
