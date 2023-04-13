package main

import (
	"fmt"

	"github.com/Masamerc/mr-garbage/garbage"
)

func main() {
	// port := os.Getenv("PORT")
	// app.Serve(port)
	weeklyScheduleRaw := garbage.ReadScheduleFromYaml()
	weeklySchedule := garbage.GetScheduleFromRawSchedule(weeklyScheduleRaw)
	fmt.Println(weeklySchedule)

}

// for local test only
// func loadDotenv() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
