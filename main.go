package main

import (
	"fmt"

	"github.com/Masamerc/mr-garbage/garbage"
)

func main() {
	// port := os.Getenv("PORT")
	// app.Serve(port)
	weeklySchedule := garbage.GetScheduleFromRawSchedule()
	fmt.Println(weeklySchedule)

}

// for local test only
// func loadDotenv() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
