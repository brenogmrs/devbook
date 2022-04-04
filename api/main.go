package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server Running!!")

	router := router.Generate()

	log.Fatal(http.ListenAndServe(":5000", router))
}
