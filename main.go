package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	//Mapping to methods
	router.ServeFiles("/static/*filepath", http.Dir("/home/sanix/tmp/static"))
	log.Fatal(http.ListenAndServe(":8000", router))
}
