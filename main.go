package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Running Hello World Handler")

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading the Body.", err)
			http.Error(rw, "Unable to read request body", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "Hello World %s", b)
	})
	log.Println("Starting the Server")
	err := http.ListenAndServe(":9090", nil)
	log.Fatal(err)
}
