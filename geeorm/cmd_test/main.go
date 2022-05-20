package main

import (
	"fmt"
	"geeorm"
)

func main() {
	engine, _ := geeorm.NewEngine("sqlite3", "gee.db")
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("drop table if exists user;").Exec()
	_, _ = s.Raw("create table user(name text)").Exec()
	_, _ = s.Raw("create table user(name text);").Exec()
	result, _ := s.Raw("insert into user(`name`) values (?), (?)", "Tom", "Jacquien").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}
