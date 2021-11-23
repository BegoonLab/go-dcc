package dcc

import (
	"log"
	"net/http"
)

func Serve() {
	log.Println("Listening on :3000...")
	fs := http.FileServer(http.Dir("../public"))
	err := http.ListenAndServe(":3000", fs)
	if err != nil {
		log.Fatal(err)
	}
}
