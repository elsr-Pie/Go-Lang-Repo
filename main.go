package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// Joke structure
type Joke struct {
	Message string `json:"message"`
}

// Madlib structure
type Madlib struct {
	Message string `json:"message"`
}

// Slice of jokes
var jokes = []string{
	"What do space cows say? Mooooooooon.",
	"What do you call a cow during an earth quake? A milk shake!",
	"What happens to an illegally parked frog? It gets toad away!",
}

// Slice of madlib components
var names = []string{"PY", "Matt", "Kia", "Shama", "Aj"}
var occupations = []string{"devOp Eng", "dcotor", "pharmacist", "politician", "actor"}
var devices = []string{"laptop", "tablet", "smartphone", "desktop", "smartwatch"}
var bodyParts = []string{"wrist", "neck", "ankle", "thigh", "shoulder"}
var moods = []string{"happy", "sad", "anxious", "excited", "angry"}
var actions = []string{"playing a calming music", "displaying motivational images", "showing a quick joke", "vibrating", "sending a life quote"}

// Function to get a random joke
func getJoke() string {
	rand.Seed(time.Now().UnixNano())
	return jokes[rand.Intn(len(jokes))]
}

// Function to get a random madlib
func getMadlib() string {
	rand.Seed(time.Now().UnixNano())
	name := names[rand.Intn(len(names))]
	age := rand.Intn(40) + 18
	occupation := occupations[rand.Intn(len(occupations))]
	device := devices[rand.Intn(len(devices))]
	bodyPart := bodyParts[rand.Intn(len(bodyParts))]
	mood := moods[rand.Intn(len(moods))]
	action := actions[rand.Intn(len(actions))]

	return name + " is a " + strconv.Itoa(age) + "-year old " + occupation + " who has been struggling with a lot of job-related stress. He/she decided to try a new application to relieve stress, which runs on a/an " + device + " to help improve his/her mood. The application senses his/her mood through a device he/she wears on his/her " + bodyPart + ". When the device senses that he/she is " + mood + ", it responds by " + action + "."
}

// JokeHandler returns a random joke
func JokeHandler(w http.ResponseWriter, r *http.Request) {
	joke := getJoke()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(joke)
}

// MadlibHandler returns a random madlib
func MadlibHandler(w http.ResponseWriter, r *http.Request) {
	madlib := getMadlib()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(madlib)
}

func main() {
	http.HandleFunc("/joke", JokeHandler)
	http.HandleFunc("/madlib", MadlibHandler)
	http.ListenAndServe(":8080", nil)
}
