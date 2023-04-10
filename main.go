package main

import (
	"os"

	"github.com/Masamerc/mr-garbage/app"
)

func main() {
	port := os.Getenv("PORT")
	app.Serve(port)
}

// for local test only
// func loadDotenv() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
