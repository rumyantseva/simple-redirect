package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	log.Print("Starting the app...")

	port := os.Getenv("PORT") // 80
	if strings.TrimSpace(port) == "" {
		log.Println("PORT environment variable has not been set")
		port = "80"
	}

	redirectTo := os.Getenv("TO")     // https://google.com
	if strings.TrimSpace(redirectTo) == "" {
		log.Println("TO environment variable has not been set")
		redirectTo = "https://google.com"
	}

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
