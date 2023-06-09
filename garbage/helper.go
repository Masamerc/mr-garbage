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
	garbages := schedule[weekday]
	if garbages == nil {
		return "No garbage collection\n"
	}
	var returnString strings.Builder
	var newLineChar string

	for index, garbage := range garbages {
		if index == len(garbages)-1 {
			newLineChar = ""
		} else {
			newLineChar = "\n\n"
		}
		returnString.WriteString(fmt.Sprintf("%s%s", garbage.FormatMessage(false), newLineChar))
	}
	return returnString.String()

}

func GetGarbageInfoFromUserMessage(userMessage string) string {
	schedule := GetScheduleFromRawSchedule()
	reverseSchedule := reverseMap(schedule)
	userMessage = strings.ToLower(strings.ReplaceAll(userMessage, " ", ""))

	switch userMessage {
	// by garbage type
	case "burnable", "general", "combustible":
		return getCollectionDays(reverseSchedule, conbustiable)
	case "plastic", "packaging":
		return getCollectionDays(reverseSchedule, plastic)
	case "cans", "bottles":
		return getCollectionDays(reverseSchedule, cansAndBottles)
	case "cardboard", "cloth":
		return getCollectionDays(reverseSchedule, CardboardAndCloth)

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
	case "week":
		return GetWeeklySchedule(schedule)
	case "tomorrow":
		weekdayTomorrow := GetTomorrowWeekDayJst()
		return getGarbageInfoResponse(weekdayTomorrow, schedule)
	case "!help":
		return helpReponse

	default:
		return fmt.Sprintf("I am sorry, I have no information regarding %s\n\n%s", strings.ToLower(userMessage), helpReponse)
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

	var schedule = make(map[string][]Garbage)
	weekdays := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	for _, weekday := range weekdays {
		garbageTypes := weeklyScheduleRaw[weekday]
		for _, garbageType := range garbageTypes {
			schedule[weekday] = append(schedule[weekday], stringToGarbage[garbageType])
		}
	}

	return schedule
}

func reverseMap(inputMap map[string][]Garbage) map[Garbage][]string {
	outputMap := make(map[Garbage][]string)

	for key, values := range inputMap {
		for _, value := range values {
			outputMap[value] = append(outputMap[value], key)
		}
	}

	return outputMap
}

func getCollectionDays(reverseSchedule map[Garbage][]string, garbage Garbage) string {
	var returnString strings.Builder
	returnString.WriteString("Collection day: \n")

	collection_days := reverseSchedule[garbage]
	for index, weekday := range collection_days {
		if index == len(collection_days)-1 {
			returnString.WriteString(fmt.Sprintf("- %s", weekday))
		} else {
			returnString.WriteString(fmt.Sprintf("- %s\n", weekday))
		}
	}
	return returnString.String()
}

func removeLastLines(s string, n int) string {
	lines := strings.Split(s, "\n")
	if len(lines) > 0 {
		lines = lines[:len(lines)-n]
	}
	return strings.Join(lines, "\n")
}

func GetWeeklySchedule(schedule map[string][]Garbage) string {
	var returnString strings.Builder
	for weekday, garbages := range schedule {
		returnString.WriteString(fmt.Sprintf("%s:\n", weekday))
		for index, garbage := range garbages {
			var newlineChar string
			if index == len(garbages)-1 { // last garbage item
				newlineChar = "\n\n"
			} else {
				newlineChar = "\n"
			}
			returnString.WriteString(fmt.Sprintf("%s%s", garbage.FormatMessage(false), newlineChar))
		}
	}
	return removeLastLines(returnString.String(), 2)
}
