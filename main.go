package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Register a handler function for the root ("/") path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Serve the "index.html" file
		http.ServeFile(w, r, "index.html")
	})

	// Serve static files (like images and JavaScript)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Specify the port to listen on
	port := 8080

	// Start the web server and listen on the specified port
	fmt.Printf("Server is listening on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
