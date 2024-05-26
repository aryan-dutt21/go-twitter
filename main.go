package main

//postgresql://twitterDB_owner:JneLFTAf3p9D@ep-small-paper-a1s81b5i.ap-southeast-1.aws.neon.tech/twitterDB?sslmode=require

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "github.com/lib/pq"
)

type App struct {
	DB *sql.DB
}

type Tweet struct {
    TweetID  int    `json:"tweetId"`
    Text     string `json:"text"`
    AuthorID int    `json:"authorId"`
}

func (app *App) setTweet(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("id")
	if userId == "" {
		http.Error(w, "Id is missing", http.StatusBadRequest)
		return
	}
	if !checkExistingUser(app.DB, userId) {
		http.Error(w, "User does not exist", http.StatusBadRequest)
		return
	}
	var tweet struct {
		Text string `json:"text"`
	}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&tweet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, queryErr := app.DB.Exec("INSERT INTO tweets (text, authorId) VALUES ($1, $2)", tweet.Text, userId)
	if queryErr != nil {
		log.Fatal(queryErr)
	}
	response := map[string]string{"message": fmt.Sprintf("Tweet added to user %s", userId)}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (app *App) getTweet(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("id")
	if !checkExistingUser(app.DB, userId) {
		http.Error(w, "User does not exist", http.StatusBadRequest)
		return
	}
	rows, err := app.DB.Query("SELECT tweetId, text, authorId FROM tweets WHERE authorId = $1", userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var tweets []Tweet
	for rows.Next() {
		var tweet Tweet
		if err := rows.Scan(&tweet.TweetID, &tweet.Text, &tweet.AuthorID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tweets = append(tweets, tweet)
	}
	response := map[string][]Tweet{"Tweets": tweets}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (app *App) createNewUser(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("id")
	var row int
	err := app.DB.QueryRow("SELECT userId FROM users WHERE userId = $1", userId).Scan(&row)
	switch {
	case err == sql.ErrNoRows:
		_, err := app.DB.Exec("INSERT INTO users (userId) VALUES ($1)", userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response := map[string]string{"response": "User created"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	case err != nil:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	default:
		response := map[string]string{"response": "User already exists"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(response)
	}
}

func main() {
	connStr := "postgresql://twitterDB_owner:JneLFTAf3p9D@ep-small-paper-a1s81b5i.ap-southeast-1.aws.neon.tech/twitterDB?sslmode=require"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	app := &App{DB: db}
	createUsersTable(db)
	createTweetsTable(db)
	defer db.Close()
	http.HandleFunc("/setTweet/{id}", app.setTweet)
	http.HandleFunc("/createNewUser/{id}", app.createNewUser)
	http.HandleFunc("/getTweet/{id}", app.getTweet)
	fmt.Println("Server listening on port 3000")
	err1 := http.ListenAndServe(":3000", nil)
	fmt.Println("Error starting server:", err1)
}

func createUsersTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS users (
		userId INTEGER UNIQUE NOT NULL,
		PRIMARY KEY(userId)
	)`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func createTweetsTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS tweets (
		tweetId SERIAL PRIMARY KEY,
		text TEXT NOT NULL,
		authorId INTEGER NOT NULL,
		FOREIGN KEY (authorId) REFERENCES users(userId)
	)`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func checkExistingUser(db *sql.DB, userId string) bool {
	var existingUserID int
	err := db.QueryRow("SELECT userId FROM users WHERE userId = $1", userId).Scan(&existingUserID)
	if err == sql.ErrNoRows {
		return false
	} else if err != nil {
		log.Printf("Error checking for existing user: %v", err)
		return false
	}
	return true
}
