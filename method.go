package main

import "fmt"

type User struct {
	id   int
	name string
}

func (self *User) TestPointer() {
	fmt.Printf("TestPointer:%p,%v\n", self, self)
}

func (self User) TestValue() {
	fmt.Printf("TestValue: %p,%v\n", &self, self)
}

func main() {
	u := User{1, "Tom"}
	fmt.Printf("User:%p,%v\n", &u, u)

	mv := User.TestValue//value copy
	mv(u)

	mp := (*User).TestValue
	mp(&u)

	mp2 := (*User).TestPointer
	mp2(&u)
}
