package main

import (
	"os"

	"github.com/Masamerc/mr-garbage/app"
)

func main() {
	port := os.Getenv("PORT")
	app.Serve(port)
	// fmt.Println(
	// 	garbage.GetGarbageInfoFromUserMessage("monday"),
	// 	garbage.GetGarbageInfoFromUserMessage("wednesday"),
	// 	garbage.GetGarbageInfoFromUserMessage("tuesday"),
	// 	garbage.GetGarbageInfoFromUserMessage("tomorrow"),
	// 	garbage.GetGarbageInfoFromUserMessage("friday"),
	// )

}

// for local test only
// func loadDotenv() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
