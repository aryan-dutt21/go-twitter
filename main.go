package main

//postgres://localhost/aryandutt?sslmode=disable

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"twitter/ent"
	"twitter/ent/user"
	tweetv1 "twitter/gen/tweet/v1"
	"twitter/gen/tweet/v1/tweetv1connect"

	"connectrpc.com/connect"
	_ "github.com/lib/pq"
)

type TweetsServer struct {
	Client *ent.Client
}

type UserServer struct {
	Client *ent.Client
}

func getUser(ctx context.Context, client *ent.Client, userId string) (*ent.User, error){
	return client.User.Query().Where(user.ID(userId)).Only(ctx)
}

func createTweet(ctx context.Context, client *ent.Client, authorId string, text string) (*ent.Tweet, error) {
	_, err := getUser(ctx, client, authorId)
	if err != nil {
		return nil, err
	}
	return client.Tweet.Create().SetAuthorID(authorId).SetID(text).SetText(text).Save(ctx)
}

func createUser(ctx context.Context, client *ent.Client, username string) (*ent.User, error) {
	return client.User.Create().SetUsername(username).SetID(username).Save(ctx)
}

func getTweets(ctx context.Context, client *ent.Client, userId string) ([]*ent.Tweet, error) {
	user, err := getUser(ctx, client, userId)
	if err != nil {
		return nil, err
	}
	return user.QueryTweets().All(ctx)
}


func (s *TweetsServer) GetTweets(
	ctx context.Context,
	req *connect.Request[tweetv1.GetTweetsRequest],
) (*connect.Response[tweetv1.GetTweetsResponse], error) {
	userId := req.Msg.UserId
	queryTweets, err := getTweets(ctx, s.Client, (userId))
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	var tweets []*tweetv1.GetTweetsResponse_Tweet
	for _, queryTweet := range queryTweets {
		tweet := &tweetv1.GetTweetsResponse_Tweet{
			AuthorId: (queryTweet.AuthorID),
			Text:     queryTweet.Text,
			TweetId:  (queryTweet.ID),
		}
		tweets = append(tweets, tweet)
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
	tweet, err := createTweet(ctx, s.Client, (userId), text)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	res := connect.NewResponse(&tweetv1.SetTweetResponse{
		Response: fmt.Sprintf("tweet %v added", tweet),
	})
	return res, nil
}

func (s *UserServer) CreateUser(
	ctx context.Context,
	req *connect.Request[tweetv1.CreateUserRequest],
) (*connect.Response[tweetv1.CreateUserResponse], error) {
	username := req.Msg.Username
	user, err := createUser(ctx, s.Client, username)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	res := connect.NewResponse(&tweetv1.CreateUserResponse{
		Response: fmt.Sprintf("user %v created", user),
	})
	return res, nil
}

func main() {
	connStr := "postgresql://twitterDB_owner:HyQCkaZJ7iM6@ep-small-paper-a1s81b5i.ap-southeast-1.aws.neon.tech/twitterDB?sslmode=require"
	client, err := ent.Open("postgres", connStr)
	if err != nil {
		log.Fatal("failed opening connection to postgres", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal("failed creating schema resources", err)
	}
	tweetsServer := &TweetsServer{client}
	userServer := &UserServer{client}
	mux := http.NewServeMux()
	mux.Handle(tweetv1connect.NewTweetsServiceHandler(tweetsServer))
	mux.Handle(tweetv1connect.NewUserServiceHandler(userServer))
	fmt.Println("Server listening on port 3000")
	err1 := http.ListenAndServe(":3000", mux)
	fmt.Println("Error starting server:", err1)
}
