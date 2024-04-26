package main

import (
	"fmt"
    "os"
	"net/http"
	"log"
	"io/ioutil"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/heathz", Heathz)
	http.HandleFunc("/secret", Secret)
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {

	name := os.Getenv("NAME")
	age  := os.Getenv("AGE")
	
	fmt.Fprint(w, "Hello, I'm %s. I'm %s.", name, age)
}

func Secret(w http.ResponseWriter, r *http.Request) {

	user := os.Getenv("USER")
	password  := os.Getenv("PASSWORD")
	
	fmt.Fprint(w, "User: %s. Password: %s", user, password)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadFile("/go/myfamily/family.txt")

	if err != nil {
		log.Fatalf("Error reading file: ", err)
	}
	
	fmt.Fprint(w, "My Family: %s", string(data))
}

// Garantir que a aplicação está saudavel
func Heathz(w http.ResponseWriter, r *http.Request) {

	duration := time.Since(startedAt)

	if duration.Seconds() < 10 || duration.Seconds() > 30 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
	} else {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}

}