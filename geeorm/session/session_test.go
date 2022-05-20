package session

import (
	"database/sql"
	"geeorm/dialect"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func NewSession() *Session {
	TestDB, _ := sql.Open("sqlite3", "../gee.db")
	dial, _ := dialect.GetDialect("sqlite3")
	return New(TestDB, dial)
}

func TestSession_Exec(t *testing.T) {
	s := NewSession()
	_, _ = s.Raw("drop table if exists user;").Exec()
	_, _ = s.Raw("create table user(name, text);").Exec()
	result, _ := s.Raw("insert into user(`name`) values (?), (?);", "Tom", "Jack").Exec()
	if count, err := result.RowsAffected(); err != nil || count != 2 {
		t.Fatal("expect 2, but got ", count)
	}
}

func TestSession_QueryRows(t *testing.T) {
	s := NewSession()
	_, _ = s.Raw("drop table if exists user;").Exec()
	_, _ = s.Raw("create table User(name, text);").Exec()
	row := s.Raw("select count(*) from user;").QueryRow()
	var count int
	if err := row.Scan(&count); err != nil || count != 0 {
		t.Fatal("Failed to query db", err)
	}
}
