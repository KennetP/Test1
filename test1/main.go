package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	//Register handlers for different paths
	http.HandleFunc("/", home)
	http.HandleFunc("/greeting", greeting)
	http.HandleFunc("/random", random)

	//Start the web server and listen for requests
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	//Parse the HTML template file
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//Define data to be used in the HTML template
	data := struct {
		Name  string
		Place string
	}{
		Name:  "Kennet Pop",
		Place: "San Pedro Columbia",
	}
	//Execute the HTML file template with the given code
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func greeting(w http.ResponseWriter, r *http.Request) {

	//Get current time and format it
	currentTime := time.Now()
	formattedTime := currentTime.Format("January 02, 2006 03:04:05 PM")

	//Write formatted time to the response
	fmt.Fprintf(w, "The current date and time is: %s", formattedTime)
}

func random(w http.ResponseWriter, r *http.Request) {
	//Define quotes
	quotes := []string{
		"Be the change you wish to see in the world-Mahatma Ganhdi",
		"Innovation distinguishes between a leader and a follower-Steve Jobs",
		"The best way to predict the future is to invent it- Alan Kay",
		"Optimism is the faith that leads to achievement. Nothing can be done without hope and confidence-Helen Keller",
		"It's not a faith in technology. It's faith in people-Steve Jobs",
	}
	//Generate a random index and select a quote
	rand.Seed(time.Now().Unix())
	randomQuote := quotes[rand.Intn(len(quotes))]
	//Write the quote to the response
	fmt.Fprintf(w, "%s", randomQuote)
}
