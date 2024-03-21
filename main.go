package main

import (
	. "REST/control"
	. "REST/dbase"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func main() {
	db, err := ConnectToDB()
	if err != nil {
		fmt.Errorf("%v", err)
	}
	defer db.Close()

	err = CreateItemsTable(db)
	if err != nil {
		fmt.Errorf("%v", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			fmt.Errorf("%v", err)
		}
		data, err := GetItems(db)
		err = tmpl.Execute(w, data)
		if err != nil {
			fmt.Errorf("%v", err)
		}
	})

	mux.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		name := r.FormValue("name")
		typee := r.FormValue("type")
		genre := r.FormValue("genre")

		item := Item{Name: name, Type: typee, Genre: genre}
		err = AddItem(item, db)
		if err != nil {
			fmt.Errorf("%v", err)
		}
		// Здесь вы можете использовать name и email для обработки данных,
		// например, сохранить в базу данных или выполнить другие действия.
		fmt.Printf("Received data: name=%s, type=%s genre=%s\n", name, typee, genre)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	mux.HandleFunc("/delete/{id}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			fmt.Errorf("%v", err)
		}
		err = DeleteItem(id, db)
		if err != nil {
			fmt.Errorf("%v", err)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Errorf("%v", err)
	}
}
