package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := os.Getenv("PORT") // 80

	redirectTo := os.Getenv("TO")     // https://google.com
	redirectType := os.Getenv("TYPE") // 301

	redirectCode, err := strconv.Atoi(redirectType)

	if err != nil {
		log.Fatal(err)
	}

	// Check of type is 301 or 302 or 303
	if redirectCode != http.StatusMovedPermanently &&
		redirectCode != http.StatusFound &&
		redirectCode != http.StatusSeeOther {
		err := fmt.Errorf("Forwarding type %s is incorrect", redirectType)
		log.Fatal(err)
	}

	log.Printf("Redirect to %s with type %d", redirectTo, redirectCode)

	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, redirectTo, redirectCode)
		},
	)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
