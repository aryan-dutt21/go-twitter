package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"twitter/ent"
	"twitter/ent/tweet"
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

func (s *TweetsServer) GetTweets(
	ctx context.Context,
	req *connect.Request[tweetv1.GetTweetsRequest],
) (*connect.Response[tweetv1.GetTweetsResponse], error) {
	userId := req.Msg.UserId
	queryTweets, err := s.Client.Tweet.Query().Where(tweet.HasAuthorWith(user.ID(int(userId)))).WithAuthor().All(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	var tweets []*tweetv1.GetTweetsResponse_Tweet
	for _, queryTweet := range queryTweets {
		tweet := &tweetv1.GetTweetsResponse_Tweet{
			AuthorId: int32(queryTweet.Edges.Author.ID),
			Text:     queryTweet.Text,
			TweetId:  int32(queryTweet.ID),
		}
		tweets = append(tweets, tweet)
	}
	return connect.NewResponse(&tweetv1.GetTweetsResponse{
		Tweets: tweets,
	}), nil
}

func (s *TweetsServer) SetTweet(
	ctx context.Context,
	req *connect.Request[tweetv1.SetTweetRequest],
) (*connect.Response[tweetv1.SetTweetResponse], error) {
	userId, text := req.Msg.UserId, req.Msg.Text
	exits, err := s.Client.User.Query().Where(user.ID(int(userId))).Exist(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	if !exits {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("user does not exist"))
	}
	tweet, err := s.Client.Tweet.Create().SetText(text).SetAuthorID(int(userId)).Save(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	return connect.NewResponse(&tweetv1.SetTweetResponse{
		Response: fmt.Sprintf("tweet %v added", tweet),
	}), nil
}

func (s *UserServer) CreateUser(
	ctx context.Context,
	req *connect.Request[tweetv1.CreateUserRequest],
) (*connect.Response[tweetv1.CreateUserResponse], error) {
	username := req.Msg.Username
	user, err := s.Client.User.Create().SetUsername(username).Save(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	return connect.NewResponse(&tweetv1.CreateUserResponse{
		Response: fmt.Sprintf("user %v created", user),
	}), nil
}

func main() {
	connStr := "postgresql://twitterDB_owner:HyQCkaZJ7iM6@ep-small-paper-a1s81b5i.ap-southeast-1.aws.neon.tech/twitterDB?sslmode=require"
	client, err := ent.Open("postgres", connStr)
	if err != nil {
		log.Fatal("failed opening connection to postgres", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal("failed creating schema resources", err)
	}
	defer client.Close()
	tweetsServer := &TweetsServer{client}
	userServer := &UserServer{client}
	mux := http.NewServeMux()
	mux.Handle(tweetv1connect.NewTweetsServiceHandler(tweetsServer))
	mux.Handle(tweetv1connect.NewUserServiceHandler(userServer))
	fmt.Println("Server listening on port 3000")
	err1 := http.ListenAndServe(":3000", mux)
	fmt.Println("Error starting server:", err1)
}
