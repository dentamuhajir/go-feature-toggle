package main

import (
	"log"
	"net/http"

	db "github.com/yasinmuhajir/go-feature-toggle/config/database"
	route "github.com/yasinmuhajir/go-feature-toggle/config/route"
)

func databaseSetUp() {
	conn, _ := db.Init()
	status, _ := db.Migration(conn)

	if status == true {
		log.Println("Succesfull added 4 migration")
	}
}

// type Articles struct {
// 	title, slug, content string
// 	isPublished          bool
// }

// func (a Articles) setTitle() string {
// }

func main() {

	go databaseSetUp()
	r := route.Init()

	log.Println(". . . server running")

	http.ListenAndServe(":8080", r)

	defer db.Close()
}
