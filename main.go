package main

import (
	"fmt"
	"log"
	"net/http"
)

var note = []string{}

func notePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "heyya (this is serving an HTML file, it's being written directly by the program. If you update this, you will have to restart to see the changes)")
}

func serveSomeHtmlPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "some.html")
}

func main() {
	http.HandleFunc("/", serveSomeHtmlPage)
	http.HandleFunc("/note", notePage)

	fmt.Printf("Starting server for getting started...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
