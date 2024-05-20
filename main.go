package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Tweet struct {
	Text string
	UserId string
}

var users = make(map[string][]Tweet)

func setTweet(w http.ResponseWriter, r *http.Request) {

	userId := r.URL.Query().Get("id")

	newTweet := Tweet{UserId: userId}

	if userId == "" {
		http.Error(w, "Id is missing", http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)

	decoder.DisallowUnknownFields()

	err := decoder.Decode(&newTweet)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users[userId] = append(users[userId], newTweet)

	response := map[string]string{"message": fmt.Sprintf("Tweet added to user %s", userId)}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}

func getTweet(w http.ResponseWriter, r *http.Request){

	userId := r.URL.Query().Get("id")

	tweets := users[userId]

	response := map[string][]Tweet{"tweets": tweets}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/setTweet", setTweet)

	http.HandleFunc("/getTweet", getTweet)

	fmt.Println("Server listening on port 3000")

	err := http.ListenAndServe(":3000", nil)

	fmt.Println("Error starting server:", err)

}
