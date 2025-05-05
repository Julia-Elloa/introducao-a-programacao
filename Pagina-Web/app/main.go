package main

import {
	"fmt"
	"log"
	"net;http"
}

func main() {
	fileserver := http.FileServer(http.Dir("./Static"))

	http.Handle("/", fileserver)

	fmt.Printf("port runningon http://localhos:8081/\n")

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
