package account

import (
	"fmt"
	"strconv"
)

type Account interface {
	Query(id string) int
	Update(id string, value int)
}

type AccountImpl struct {
	Name  string
	Id    string
	Value int
}

func (a *AccountImpl) Query(_ string) int {
	fmt.Println("AccountImpl.Query")
	return 100
}

func (a *AccountImpl) Update(_ string, num int) {
	fmt.Println("AccountImpl.Update")
	atoi, _ := strconv.Atoi(a.Id)
	temp := atoi + num
	itoa := strconv.Itoa(temp)
	a.Id = itoa
	fmt.Println("After update :", a.Id)
}

func New(name, id string) Account {
	a := &AccountImpl{name, id, 100}
	return a
}
