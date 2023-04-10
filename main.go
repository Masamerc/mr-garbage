package main

import (
	"os"

	"github.com/Masamerc/mr-garbage/server"
)

func main() {
	port := os.Getenv("PORT")
	server.Serve(port)
}
