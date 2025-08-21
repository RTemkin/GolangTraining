package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name       string
	Age        uint16
	Money      int
	Avg_grades float64
	Happiness  float64
}

func (u *User) GetAllInfo() string {
	return fmt.Sprintf("User name is: %s. Hi is: %d fnd hi has money equal: %d\n", u.Name, u.Age, u.Money)
}

func (u *User) SetNewName(newname string) {
	u.Name = newname
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{
		Name:       "Bob",
		Age:        25,
		Money:      -3,
		Avg_grades: 4.2,
		Happiness:  0.8,
	}
	// fmt.Fprintf(w, "User name is: %v\n", bob.Name)

	// bob.SetNewName("Alex")
	// fmt.Fprintf(w, bob.GetAllInfo())

	// fmt.Fprintf(w, "<b>Main text<b>")

	// подгрузка шаблона html на страницу
	templ, _ := template.ParseFiles("templates/home_page.html")
	templ.Execute(w, bob)

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

	handelRequest()

}
