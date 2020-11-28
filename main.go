package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = 3000

func main() {
	router := ChiRouter().InitRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), router))
}
