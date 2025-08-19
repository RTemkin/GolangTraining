package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page.")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем метод запроса
	if r.Method == http.MethodGet {
		// Отображаем форму
		fmt.Fprintf(w, `
            <form method="POST">
                <input type="text" name="name" placeholder="Enter your name">
                <button type="submit">Submit</button>
            </form>
        `)
	} else if r.Method == http.MethodPost {
		// Обрабатываем данные формы
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(w, "Error parsing form: %v", err)
			return
		}

		name := r.FormValue("name")
		fmt.Fprintf(w, "Hello, %s!", name)
	}
}

// создаем сервер на порту 8080 который при любом запросе будет отdечать Hello World
func main() {

	// Регистрируем обработчик для всех запросов
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/form", formHandler)

	// Запускаем сервер на порту 8080
	fmt.Println("Starting server at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}

}
