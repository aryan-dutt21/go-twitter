package main

//postgres://localhost/aryandutt?sslmode=disable

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	tweetv1 "twitter/gen/tweet/v1"
	"twitter/gen/tweet/v1/tweetv1connect"

	"connectrpc.com/connect"
	_ "github.com/lib/pq"
)

type TweetsServer struct {
	DB *sql.DB
}

type UserServer struct {
	DB *sql.DB
}

func (s *TweetsServer) GetTweets(
	ctx context.Context,
	req *connect.Request[tweetv1.GetTweetsRequest],
) (*connect.Response[tweetv1.GetTweetsResponse], error) {
	userId := req.Msg.UserId
	if userId == 0 {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("userId is required"))
	}
	rows, err := s.DB.Query("SELECT tweetId, text, authorId FROM tweets WHERE authorId = $1", userId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	defer rows.Close()
	var tweets []*tweetv1.GetTweetsResponse_Tweet
	for rows.Next() {
		var tweet tweetv1.GetTweetsResponse_Tweet
		if err := rows.Scan(&tweet.TweetId, &tweet.Text, &tweet.AuthorId); err != nil {
			return nil, connect.NewError(connect.CodeInternal, err)
		}
		tweets = append(tweets, &tweet)
	}
	if err := rows.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	res := connect.NewResponse(&tweetv1.GetTweetsResponse{
		Tweets: tweets,
	})
	return res, nil
}

func (s *TweetsServer) SetTweet(
	ctx context.Context,
	req *connect.Request[tweetv1.SetTweetRequest],
) (*connect.Response[tweetv1.SetTweetResponse], error) {
	userId, text := req.Msg.UserId, req.Msg.Text
	if userId == 0 {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("userId is required"))
	}
	if text == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("text is required"))
	}
	if !checkExistingUser(s.DB, userId) {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("user does not exist"))
	}
	_, err := s.DB.Exec("INSERT INTO tweets (text, authorId) VALUES ($1, $2)", text, userId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	res := connect.NewResponse(&tweetv1.SetTweetResponse{
		Response: fmt.Sprintf("tweet added to user %v", userId),
	})
	return res, nil
}

func (s *UserServer) CreateUser(
	ctx context.Context,
	req *connect.Request[tweetv1.CreateUserRequest],
) (*connect.Response[tweetv1.CreateUserResponse], error) {
	username := req.Msg.Username
	if username == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("username is required"))
	}
	_, err := s.DB.Exec("INSERT INTO users (username) VALUES ($1)", username)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	res := connect.NewResponse(&tweetv1.CreateUserResponse{
		Response: fmt.Sprintf("user %v created", username),
	})
	return res, nil
}

func main() {
	connStr := "postgresql://twitterDB_owner:HyQCkaZJ7iM6@ep-small-paper-a1s81b5i.ap-southeast-1.aws.neon.tech/twitterDB?sslmode=require"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to DB", err)
	}
	mux := http.NewServeMux()
	defer db.Close()
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
	tweetsServer := &TweetsServer{db}
	userServer := &UserServer{db}
	mux.Handle(tweetv1connect.NewTweetsServiceHandler(tweetsServer))
	mux.Handle(tweetv1connect.NewUserServiceHandler(userServer))
	fmt.Println("Server listening on port 3000")
	err1 := http.ListenAndServe(":3000", mux)
	fmt.Println("Error starting server:", err1)
}

func createUsersTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS users (
		userId SERIAL PRIMARY KEY,
		username TEXT NOT NULL
	)`
	_, err := db.Exec(query)
	return err
}

func createTweetsTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS tweets (
		tweetId SERIAL PRIMARY KEY,
		text TEXT NOT NULL,
		authorId INTEGER NOT NULL,
		FOREIGN KEY (authorId) REFERENCES users(userId)
	)`
	_, err := db.Exec(query)
	return err
}

func checkExistingUser(db *sql.DB, userId int32) bool {
	var existingUsername string
	err := db.QueryRow("SELECT username FROM users WHERE userId = $1", userId).Scan(&existingUsername)
	if err == sql.ErrNoRows {
		return false
	} else if err != nil {
		log.Printf("Error checking for existing user: %v", err)
		return false
	}
	return true
}
