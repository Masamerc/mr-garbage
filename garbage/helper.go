package garbage

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v2"
)

var helpReponse string = fmt.Sprintf(`To get information from me, you need to provide weekday or garbage type. here are some examples.
garbage type %s "burnable", "general", "combustible", "cans", "bottles", "plastic"
weekday %s "Monday", "Tuesday", "Friday"`, rightArrow, rightArrow)

func GetGarbageInfoFromUserMessage(userMessage string) string {
	userMessage = strings.ToLower(strings.ReplaceAll(userMessage, " ", ""))
	switch userMessage {
	// by garbage type
	case "burnable", "general", "combustible":
		return "Collection day: Tuesday & Saturday"
	case "plastic", "packaging":
		return "Collection day: Monday"
	case "cans", "bottles":
		return "Collection day: Friday"
	case "cardboard", "cloth":
		return "Collection day: Friday"
	// by weekday
	case "monday":
		return plastic.FormatMessage(false)
	case "tuesday":
		return conbustiable.FormatMessage(false)
	case "wednesday", "thursday", "sunday":
		return "No garbage collection"
	case "friday":
		return fmt.Sprintf("%s\n%s",
			cansAndBottles.FormatMessage(false),
			CardboardAndCloth.FormatMessage(false),
		)
	case "saturday":
		return conbustiable.FormatMessage(false)
	// special command
	case "tomorrow":
		weekdayTomorrow := GetTomorrowWeekDayJst()
		return Schedule[weekdayTomorrow].FormatMessage(true)
	case "!help":
		return helpReponse

	default:
		return fmt.Sprintf("I am sorry, I have no information regarding %s", strings.ToLower(userMessage))
	}
}

func ReadScheduleFromYaml() map[string][]string {
	file, err := ioutil.ReadFile("schedule.yaml")
	if err != nil {
		log.Fatal(err)
	}

	data := make(map[string]map[string][]string)

	err = yaml.Unmarshal(file, &data)
	if err != nil {
		log.Fatal((err))
	}

	return data["weekly_schedule"]
}

var stringToGarbage = map[string]Garbage{
	"burnable":  conbustiable,
	"plastic":   plastic,
	"cans":      cansAndBottles,
	"cardboard": CardboardAndCloth,
}

func GetScheduleFromRawSchedule(weeklyScheduleRaw map[string][]string) map[string][]Garbage {
	var Schedule = make(map[string][]Garbage)
	weekdays := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	for _, weekday := range weekdays {
		garbageTypes := weeklyScheduleRaw[weekday]
		for _, garbageType := range garbageTypes {
			Schedule[weekday] = append(Schedule[weekday], stringToGarbage[garbageType])
		}
	}

	return Schedule
}
