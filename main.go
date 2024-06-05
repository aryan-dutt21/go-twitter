package main

//postgres://localhost/aryandutt?sslmode=disable

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
	AuthorName string    `json:"authorUsername"`
}

type Response struct {
	Message string `json:"message"`
}

type TweetsResponse struct {
	Data []Tweet `json:"data"`
}

type CreateUserRequest struct {
	Username    string `json:"username"`
	Displayname string `json:"displayname"`
}

func (app *App) setTweet(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	if username == "" {
		http.Error(w, "username is required", http.StatusBadRequest)
		return
	}
	if !checkExistingUser(app.DB, username) {
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, queryErr := app.DB.Exec("INSERT INTO tweets (text, authorUsername) VALUES ($1, $2)", tweet.Text, username)
	if queryErr != nil {
		http.Error(w, "Error adding tweet", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{Message: fmt.Sprintf("Tweet added to user %s", username)})
}

func (app *App) getTweets(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	if username == "" {
		http.Error(w, "username is required", http.StatusBadRequest)
		return
	}
	rows, err := app.DB.Query("SELECT tweetId, text, authorUsername FROM tweets WHERE authorUsername = $1", username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	
	var tweets []Tweet
	for rows.Next() {
		var tweet Tweet
		if err := rows.Scan(&tweet.TweetID, &tweet.Text, &tweet.AuthorName); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tweets = append(tweets, tweet)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(TweetsResponse{Data: tweets})
}

func (app *App) createUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if req.Username == "" {
		http.Error(w, "username is required", http.StatusBadRequest)
		return
	}
	if req.Displayname == "" {
		http.Error(w, "display name is required", http.StatusBadRequest)
		return
	}
	_, err = app.DB.Exec("INSERT INTO users (username, displayname) VALUES ($1, $2)", req.Username, req.Displayname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Message: fmt.Sprintf("User %s created", req.Username)})
}

//getUserInfo

func main() {
	connStr := "postgresql://twitterDB_owner:HyQCkaZJ7iM6@ep-small-paper-a1s81b5i.ap-southeast-1.aws.neon.tech/twitterDB?sslmode=require"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to DB", err)
	}
	mux := http.NewServeMux()
	defer db.Close()
	app := &App{DB: db}
	err = createUsersTable(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = createTweetsTable(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	mux.HandleFunc("/setTweet/{username}", app.setTweet)
	mux.HandleFunc("/createUser", app.createUser)
	mux.HandleFunc("/getTweets/{username}", app.getTweets)
	fmt.Println("Server listening on port 3000")
	err1 := http.ListenAndServe(":3000", mux)
	fmt.Println("Error starting server:", err1)
}

func createUsersTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS users (
		username TEXT PRIMARY KEY NOT NULL,
		displayname TEXT NOT NULL
	)`
	_, err := db.Exec(query)
	return err
}

func createTweetsTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS tweets (
		tweetId SERIAL PRIMARY KEY,
		text TEXT NOT NULL,
		authorUsername TEXT NOT NULL,
		FOREIGN KEY (authorUsername) REFERENCES users(username)
	)`
	_, err := db.Exec(query)
	return err
}

func checkExistingUser(db *sql.DB, username string) bool {
	var existingUsername string
	err := db.QueryRow("SELECT username FROM users WHERE username = $1", username).Scan(&existingUsername)
	if err == sql.ErrNoRows {
		return false
	} else if err != nil {
		log.Printf("Error checking for existing user: %v", err)
		return false
	}
	return true
}
