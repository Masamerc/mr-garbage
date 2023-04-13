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

var stringToGarbage = map[string]Garbage{
	"burnable":  conbustiable,
	"plastic":   plastic,
	"cans":      cansAndBottles,
	"cardboard": CardboardAndCloth,
}

func getGarbageInfoResponse(weekday string, schedule map[string][]Garbage) string {
	if garbages := schedule[weekday]; garbages == nil {
		return "No garbage collection\n"
	} else {
		var returnString string
		for _, garbage := range garbages {
			returnString += garbage.FormatMessage(false) + "\n\n"
		}
		return returnString
	}
}

func GetGarbageInfoFromUserMessage(userMessage string) string {
	schedule := GetScheduleFromRawSchedule()
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
		return getGarbageInfoResponse("Mon", schedule)
	case "tuesday":
		return getGarbageInfoResponse("Tue", schedule)
	case "wednesday":
		return getGarbageInfoResponse("Wed", schedule)
	case "thursday":
		return getGarbageInfoResponse("Thu", schedule)
	case "friday":
		return getGarbageInfoResponse("Fri", schedule)
	case "saturday":
		return getGarbageInfoResponse("Sat", schedule)
	case "suncday":
		return getGarbageInfoResponse("Sun", schedule)

	// special command
	case "tomorrow":
		weekdayTomorrow := GetTomorrowWeekDayJst()
		return getGarbageInfoResponse(weekdayTomorrow, schedule)
	case "!help":
		return helpReponse

	default:
		return fmt.Sprintf("I am sorry, I have no information regarding %s", strings.ToLower(userMessage))
	}
}

func ReadRawScheduleFromYaml() map[string][]string {
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

func GetScheduleFromRawSchedule() map[string][]Garbage {
	weeklyScheduleRaw := ReadRawScheduleFromYaml()

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
