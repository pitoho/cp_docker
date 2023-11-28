package main

import (
 "fmt"
 "log"
 "net/http"
 "github.com/go-redis/redis/v8"
 "github.com/gorilla/mux"
)

var redisClient *redis.Client

func init() {
 	redisClient = redis.NewClient(&redis.Options{
  	Addr:     "redis:6379", 
  	Password: "",           
  	DB:       0,            
 	})
}

func handler(w http.ResponseWriter, r *http.Request) {
 	key := "example-key"
 	err := redisClient.Set(r.Context(), key, "Hello, Redis!", 0).Err()
 	if err != nil {
  		http.Error(w, "Error writing to Redis", http.StatusInternalServerError)
  		return
 	}

 	val, err := redisClient.Get(r.Context(), key).Result()
 	if err != nil {
  		http.Error(w, "Error reading from Redis", http.StatusInternalServerError)
  		return
 	}

 	fmt.Fprintf(w, "Data from Redis: %s", val)
}

func main() {
 	r := mux.NewRouter()
 	r.HandleFunc("/", handler)

 	port := "8080"
 	log.Printf("Server listening on :%s", port)
 	log.Fatal(http.ListenAndServe(":"+port, r))
}