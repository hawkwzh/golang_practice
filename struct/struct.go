package main

import "fmt"

func main() {
	createStruct()
}

//User 创建struct
type User struct {
	name   string
	age    int
	salary float64
}

func createStruct() {
	var user = User{"wu.zhenhuan", 22, 123.4}
	fmt.Printf("%v", user)
	fmt.Printf("\n name: %v \n age: %v \n salary: %v", user.name, user.age, user.salary)
}
