package main

import (
	"fmt"
    "os"
	"net/http"
	"log"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {

	name := os.Getenv("NAME")
	age  := os.Getenv("AGE")
	
	fmt.Fprint(w, "Hello, I'm %s. I'm %s", name, age)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadFile("/go/myfamily/family.txt")

	if err != nil {
		log.Fatalf("Error reading file: ", err)
	}
	
	fmt.Fprint(w, "My Family: %s", string(data))
}