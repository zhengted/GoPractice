package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func testRedis() bool {
	// e.g. REDIS_PORT = tcp://172.17.0.89:6379
	conn, err := redis.DialTimeout("0.0.0.0", "6379", 0, 1*time.Second, 1*time.Second)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer conn.Close()

	size, err := conn.Do("DBSIZE")
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Printf("DB size is %d \n", size)

	_, err = conn.Do("SET", "user:user0", 123)
	_, err = conn.Do("SET", "user:user1", 456)
	_, err = conn.Do("APPEND", "user:user0", 87)

	user0, err := redis.Int(conn.Do("GET", "user:user0"))
	user1, err := redis.Int(conn.Do("GET", "user:user1"))

	fmt.Printf("user0 is %d , user1 is %d \n", user0, user1)
	return true
}

func main() {
	fmt.Println(time.Now(), "[Redis test start]")

	if testRedis() {
		fmt.Println("[Redis test ok]")
	} else {
		fmt.Println("[Redis test failed]")
	}
}
func BuildTree103() *TreeNode {
	root := TreeNode{Val: 3}
	root.Left = &TreeNode{9, nil, nil}
	root.Right = &TreeNode{20, nil, nil}
	root.Right.Left = &TreeNode{15, nil, nil}
	root.Right.Right = &TreeNode{7, nil, nil}
	return &root
}
func groupAnagramsDemo() {
	param := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagramsEx(param))
}
