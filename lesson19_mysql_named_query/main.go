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

func selectNamedQuery() {
	sql := "select * from student where age = :age"
	rows, err := db.NamedQuery(sql, map[string]interface{}{
		"age": 19,
	})

	if err != nil {
		fmt.Printf("selectNamedQuery failed, err: %v \n", err)
	}

	defer rows.Close()

	for rows.Next() {
		var u User
		if err := rows.StructScan(&u); err != nil {
			fmt.Printf("StructScan failed, err: %v \n", err)
			continue
		}
		fmt.Printf("%v \n", u)
	}
}

// batchInsert 批量插入数据
func batchInsert() {
	users := []User{{Name: "周", Age: 12}, {Name: "吴", Age: 16}}

	sql := "insert into student (name, age) values (:name, :age)"
	_, err := db.NamedExec(sql, users)

	if err != nil {
		fmt.Printf("batchInsert failed: %v \n", err)
		return
	}
	fmt.Println("batch insert success")
}

func main() {
	if err := initializeDatabase(); err != nil {
		panic(err)
	}

	fmt.Println("connect to db successful")

	//selectNamedQuery()

	batchInsert()
}
