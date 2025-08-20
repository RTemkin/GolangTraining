package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name       string
	age        uint16
	money      int
	avg_grades float64
	happiness  float64
}

func (u User) GetAllInfo() string {
	return fmt.Sprintf("User name is: %s. Hi is: %d fnd hi has money equal: %d", u.name, u.age, u.money)
}

func(u *User)SetNewName(newname string){
	u.name = newname
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{
		name:       "Bob",
		age:        25,
		money:      -3,
		avg_grades: 4.2,
		happiness:  0.8,
	}
	fmt.Fprintf(w, "User name is: %v\n", bob.name)

	bob.SetNewName("Alex")
	fmt.Fprintf(w, bob.GetAllInfo())

}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts page!")
}

func handelRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)

	fmt.Println("Starting custom server at port 8080")
	http.ListenAndServe(":8080", nil)
}

func main() {

	// bob := User{
	// 	name: "Bob",
	// 	age: 25,
	// 	money: -3,
	// 	avg_grades: 4.2,
	// 	happiness: 0.8,
	// }

	handelRequest()

}
