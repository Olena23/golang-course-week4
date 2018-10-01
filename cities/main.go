package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"time"
	"fmt"
	"os"
	"encoding/csv"
	"bufio"
	"io"
	"log"
	"strings"
	"math/big"
	"crypto/rand"
)

var cities []string
var response = ""
var newResponse = ""
//to distinguish from checking last letter of submitted answer
var response1 = ""
var newResponse1 = ""
var givenCity = ""
var goOn = true


func main() {

	//Form list of city names


	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Here's a game of cities"))
	})

	r.Route("/", func(r chi.Router) {
		r.Get("/init", startGame)
		r.Get("/get-answer/{city_name}", getAnswer)
		r.Post("/submit/{answer}", submit)

	})
	http.ListenAndServe(":3333", r)
}

func startGame (w http.ResponseWriter, r *http.Request) {
	csvFile, _ := os.Open("cities.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		addName := strings.Trim(line[1], " ")
		cities = append(cities, addName)
	}

	//delete header from cities
	cities = append(cities[:0], cities[1:]...)
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(cities))))
	if err != nil {
		panic(err)
	}
	n := nBig.Int64()
	givenCity = cities[n]
	fmt.Printf("Here is a random %T in [0,27) : %d\n", n, n)
	fmt.Fprint(w, "City: " + givenCity + "\n")
	//delete used city
	cities = append(cities[:n], cities[n+1:]...)
	return
}

func getAnswer (w http.ResponseWriter, r *http.Request) {
	cityName := chi.URLParam(r, "city_name")
	checkLastLetter(cityName, cities)
	fmt.Fprint(w, "City: " + response + "\n")
	return
}

func submit (w http.ResponseWriter, r *http.Request) {
	answer := chi.URLParam(r, "answer")
	for j :=1; j < 6 ; j++ {
		if (len(givenCity)-j>=0) && ((len(givenCity)-j+1)>=0) {
			var lastLetter = givenCity[len(givenCity)-j:len(givenCity)-j+1]
				if answer[:1] == lastLetter {
					newResponse1 = answer
					if response1 != newResponse1{
						response1 = newResponse1
						w.WriteHeader(200)
						givenCity = answer
						//NEVER WORKS
						fmt.Fprintf(w, "Your answer is correct: %q \n", answer)
						checkLastLetter(response, cities)
						fmt.Fprint(w, "New city: " + response + "\n")
						return
					}
				}
			}
		}
	return
}


func checkLastLetter (s string, cities []string) {
	for j :=1; j < 6 ; j++ {
		if (len(s)-j>=0) && ((len(s)-j+1)>=0) {
			var lastLetter = s[len(s)-j:len(s)-j+1]
			for i := 0; i < len(cities); i++ {
				if cities[i][:1] == lastLetter {
					newResponse = cities[i]
					cities = append(cities[:i], cities[(i+1):]...)
					if response != newResponse{
						response = newResponse
						return
					}
				}
			}
		}
	}
	goOn = false
}
