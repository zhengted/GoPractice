package main

import (
	"GoPractice/lang/Interceptor/account"
	_ "GoPractice/lang/Interceptor/proxy"
)

func main() {
	id := "100111"
	a := account.New("hu", id)
	a.Query(id)
	a.Update(id, 500)
}
