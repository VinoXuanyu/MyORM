package session

import (
	"geeorm/log"
	"testing"
)

type Account struct {
	ID       int `geeorm:"primary key""`
	Password string
}

func (account *Account) AfterQuery(s *Session) error {
	log.Info("after query", account)
	account.Password = "******"
	return nil
}

func TestSession_CallMethod(t *testing.T) {
	s := NewSession().Model(&Account{})
	_ = s.DropTable()
	_ = s.CreateTable()
	_, _ = s.Insert(&Account{1, "123456"})
	u := &Account{}
	err := s.First(u)
	if err != nil || u.Password != "******" {
		t.Fatal("Failed to call method")
	}

}
