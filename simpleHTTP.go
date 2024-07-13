package main

import (
	"fmt"
	"log"
	"os"
 	"net/http"
)

func main() {
	log.Print("simplehttp: Enter main()")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

// printing request headers/params
func handler(w http.ResponseWriter, r *http.Request) {

	log.Print("request from address: %q\n", r.RemoteAddr)
	fmt.Fprintf(w, "Env message = %s\n", os.Getenv("message"))
}
