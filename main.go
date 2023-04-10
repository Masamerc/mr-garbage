package main

import (
	"os"

	"github.com/Masamerc/mr-garbage/server"
)

func main() {
	port := os.Getenv("PORT")
	server.Serve(port)
}

// for local test only
// func loadDotenv() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
