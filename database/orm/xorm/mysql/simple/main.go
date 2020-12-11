package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func main() {

	engine, _ := xorm.NewEngine("mysql", "root:1997@tcp(127.0.0.1:3306)/todo?charset=utf8")
	sql := "SELECT * FROM user"
	//results, err := engine.Query(sql)
	results, err := engine.QueryInterface(sql)
	//results, err := engine.QueryString(sql)
	if err != nil {
		log.Fatal("query", sql, err)
		return
	}
	total := len(results)
	if total == 0 {
		fmt.Println("没有任何数据", sql)
	} else {
		for i, data := range results {
			fmt.Printf("%d = %v\n", i, data)
		}
	}

}
