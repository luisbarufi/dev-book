package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	config.LoadEnv()
	utils.RenderTemplates()
	r := router.Generate()

	fmt.Printf("Webapp running on port %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
