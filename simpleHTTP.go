package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("simplehttp: Enter main()")
	http.HandleFunc("/", handler)
	http.ListenAndServe("0.0.0.0:8000", nil)
}

// printing request headers/params
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", os.Getenv("message"))
}
