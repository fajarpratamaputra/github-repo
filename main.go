package main

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"net/http"
	"fmt"
)


func repo(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("https://api.github.com/users/fajarpratamaputra/repos?go&sort=stars&order=desc")

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		log.Fatal("Unexpected status code", res.StatusCode)
	}

	fmt.Printf("Body: %s\n", body)
	return
}

func main() {
	http.HandleFunc("/repo", repo)

    fmt.Println("starting web server at http://localhost:8080/")
    http.ListenAndServe(":8080", nil)	
}

