package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
)

func main() {
	fmt.Println("Web app running")

	r := router.Generate()
	log.Fatal(http.ListenAndServe(":3000", r))
}
