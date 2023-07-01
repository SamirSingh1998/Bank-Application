package app

import (
	"log"
	"net/http"
)

func Start() {
	//define routes
	http.HandleFunc("/greet", greet)

	http.HandleFunc("/customers", getAllCustomer)

	//starting server
	//Listen and Serve starts the server at localhost:port and second parameter is handler which is a request multiplexer, and since we are relying on default so we will keep it as nil.
	//http.ListenAndServe("localhost:8000", nil) //To ensure that your program runs continuously, you can run the http.ListenAndServe function call in a separate goroutine.
	go func() {
		if err := http.ListenAndServe("localhost:8000", nil); err != nil {
			log.Fatal(err)
		}
	}()
	log.Println("Server is up and running on localhost:8000")

	// after running main. go
	// curl http://localhost:9000/greet
	//	Hello World

	select {} // Use a blocking mechanism to keep the program running
}
