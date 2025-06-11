package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/IkhwanAL/job-tracker/dbService"
	"github.com/IkhwanAL/job-tracker/utility/cryptography"
	"github.com/IkhwanAL/job-tracker/views"
	"github.com/joho/godotenv"
)

func main() {
	// runSeed()

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
			//TODO Add Something To Trigger 404 or Something Error Show up
			fmt.Printf("username not exists %s", err.Error())
			w.Header().Set("HX-Swap", "none")
			return
		}

		err = cryptography.Compare([]byte(password), []byte(result.Password))

		if err != nil {
			fmt.Print("password is not the same")
			w.Header().Set("HX-Swap", "none")
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
