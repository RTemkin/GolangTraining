package main

import "fmt"

type User struct {
	Name string
	Age  int
	Prof string
	Iq   float32
}

func CreatUser(name string, age int, prof string, iq float32)User {
	return User{
		Name: name,
		Age:  age,
		Prof: prof,
		Iq:   iq,
	}
}

func (u *User)ismIq(iq float32)User{
	return User{
		Name: u.Name,
		Age:  u.Age,
		Prof: u.Prof,
		Iq:   iq,
	}
}

func (u *User)changeIq(iq float32){
	u.Iq = iq
}


func main() {

		
	user1 := CreatUser("Rom", 33, "Ingener", 100.5)

	fmt.Println(user1)

	user1.changeIq(120)
	user2 := user1.ismIq(110.4)

	fmt.Println(user1)
	fmt.Println(user2)

}
