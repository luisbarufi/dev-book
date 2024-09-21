package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	utils.RenderTemplates()
	r := router.Generate()

	fmt.Println("Webapp running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
