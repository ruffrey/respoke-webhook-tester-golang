package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if strings.ToLower(r.Method) != "post" {
		log.Println("Unexpected Request:", r.Method, r.URL.String())
		w.WriteHeader(420)
		w.Write([]byte("Enhance your calm"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Request Error Reading:", r.Method, r.URL.String(), "  ", err)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	log.Println("Request:", r.Method, r.URL.String())
	log.Println("  ", string(body))
	w.WriteHeader(200)
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal("First argument should be port")
	}
	port := args[0]

	if port == "" {
		log.Fatal("First argument should be port")
	}

	http.HandleFunc("/", webhookHandler)

	log.Println("Starting server on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
