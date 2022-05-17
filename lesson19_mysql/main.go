package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:Hunan19970708!@/offline_gin_demo")

	if err != nil {
		panic(err)
	}

	// 尝试建立数据库连接
	if err := db.Ping(); err != nil {
		fmt.Println("连接数据库失败...")
		panic(err)
	}
	// func (db *DB) SetMaxOpenConns(n int) 设置与数据库建立连接的最大数目。如果n大于0且小于最大闲置连接数，会将最大闲置连接数减小到匹配最大开启连接数的限制。
	// 如果 n<=0，不会限制最大开启连接数，默认为0（无限制）
	db.SetMaxOpenConns(10)

	// func (db *DB) SetMaxIdleConns(n int) 设置连接池中最大闲置连接数。
	// 如果n大于最大开启连接数，则新的最大闲置连接数会减小到匹配最大开启连接数的限制。
	// 如果 n<=0，不会保留闲置连接
	db.SetMaxIdleConns(10)
	fmt.Println("连接数据库成功...")
}
