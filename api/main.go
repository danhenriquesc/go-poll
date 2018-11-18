package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "github.com/go-redis/redis"
)

type PollResponse struct {
    Candidate1 string `json:"candidate_1"`
    Candidate2 string `json:"candidate_2"`
}

var RedisClientSingleton *redis.Client

const GET = "GET"
const POST = "POST"

func GetCandidateVotes(candidate string) string {
    client := RedisClient()
    candidate_votes, err := client.Get(candidate).Result()
    if err != nil {
        panic(err)
    }

    return candidate_votes
}

func GetCandidatesVotes(candidates ...string) []string {
    var candidates_votes []string

    client := RedisClient()
    
    candidates_votes_interfaces, err := client.MGet(candidates...).Result()

    if err != nil {
        panic(err)
    }

    for _, candidates_votes_interface := range candidates_votes_interfaces {
        candidate_votes, ok := candidates_votes_interface.(string)
        if ok == false {
            candidate_votes = "0"
        }
        candidates_votes = append(candidates_votes, candidate_votes)
    }

    return candidates_votes
}

func SetVote(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    candidate := params["candidate"]

    client := RedisClient()
    err := client.Incr(candidate).Err()
    if err != nil {
        panic(err)
    }

    candidates_votes := GetCandidatesVotes("1", "2")

    w.WriteHeader(200)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(&PollResponse{candidates_votes[0], candidates_votes[1]})
}

func GetVotes(w http.ResponseWriter, r *http.Request) {
    candidates_votes := GetCandidatesVotes("1", "2")

    w.WriteHeader(200)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(&PollResponse{candidates_votes[0], candidates_votes[1]})
}

func RedisClient() *redis.Client{
    if RedisClientSingleton == nil {
        RedisClientSingleton = redis.NewClient(&redis.Options{
            Addr:     "redis:6379",
            Password: "",
            DB:       0,
        })
    }
    
    return RedisClientSingleton
}

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/vote", GetVotes).Methods(GET)
    router.HandleFunc("/vote/{candidate:[1|2]}", SetVote).Methods(POST)
    log.Fatal(http.ListenAndServe(":8000", router))
}