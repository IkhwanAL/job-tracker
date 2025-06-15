package main

// Abandon all hope, ye who code here

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/IkhwanAL/job-tracker/dbService"
	"github.com/IkhwanAL/job-tracker/utility/cryptography"
	"github.com/IkhwanAL/job-tracker/views"
	"github.com/IkhwanAL/job-tracker/views/components"
	"github.com/joho/godotenv"
)

func main() {
	// tmp.RunSeed()

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := sql.Open("sqlite3", "file:job_tracker.db")

	if err != nil {
		log.Fatalf("error connection db: %s", err.Error())
	}

	defer db.Close()

	queries := dbService.New(db)

	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := views.Login().Render(r.Context(), w)
		if err != nil {
			log.Fatal(err.Error())
		}
	})

	mux.HandleFunc("POST /login", func(w http.ResponseWriter, r *http.Request) {

		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		result, err := queries.GetUser(r.Context(), sql.NullString{String: username, Valid: true})

		if err != nil {
			w.Header().Set("Content-Type", "text/html")
			w.Header().Set("HX-Retarget", "#error-container")
			w.Header().Set("HX-Reswap", "innerHTML")
			err := components.Popup("username does not exists", "error").Render(r.Context(), w)
			if err != nil {
				log.Fatal(err.Error())
			}
			return
		}
		t1 := time.Now()
		err = cryptography.Compare([]byte(password), []byte(result.Password))
		t2 := time.Now()

		fmt.Println(16, ": ", t2.Sub(t1))

		if err != nil {
			w.Header().Set("Content-Type", "text/html")
			w.Header().Set("HX-Retarget", "#error-container")
			w.Header().Set("HX-Reswap", "innerHTML")
			err := components.Popup("password does not match", "error").Render(r.Context(), w)
			if err != nil {
				log.Fatal(err.Error())
			}
			return
		}

		err = views.Index().Render(r.Context(), w)

		if err != nil {
			log.Fatal(err.Error())
		}
	})

	err = http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal(err.Error())
	}
}
