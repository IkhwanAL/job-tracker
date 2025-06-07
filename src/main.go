package main

import (
	"log"
	"net/http"

	"github.com/IkhwanAL/job-tracker/views"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := views.Index().Render(r.Context(), w)

		if err != nil {
			log.Fatal(err.Error())
		}
	})

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err.Error())
	}
}
