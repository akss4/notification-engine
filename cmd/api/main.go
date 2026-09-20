package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { // ww and r is wrtite and read means we define a handler which when pinged runs this function which params two value
		w.Header().Set("Content-Type", "application/json") //tells  client that the responce is json
		fmt.Fprintln(w,
			`{
		"status": "ok"
		}`)
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return // stop the execution of the handler and return a response to the client
		} // only woeks if get == true
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w,
			`{
		"message": "Hello, World!"
		}`)
	})

	fmt.Println("Starting server on :8080")

	err := http.ListenAndServe(":8080", nil) //listen and serve wait for network connection and handles incoming requests
	if err != nil {                          // it also says use port 8080
		fmt.Println("Error starting server:", err)
	}

}
