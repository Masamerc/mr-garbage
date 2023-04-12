package app

import (
	"fmt"
	"strings"
)

const rightArrow string = "\u27A1"

type Garbage struct {
	Type    string
	Details string
	Emoji   string
}

func (g Garbage) FormatMessage(prefixed bool) string {
	var prefix string
	if prefixed {
		prefix = "Garbage day tomorrow: "
	}

	if g.Type == "" {
		g.Type = "No Garbage"
		g.Details = "no garbage collection."
	}

	messageText := fmt.Sprintf("%s%s %s\n%s %s", prefix, g.Type, g.Emoji, rightArrow, g.Details)
	return messageText
}

// garbage type definitions
var conbustiable Garbage = Garbage{
	Type:    "burnable",
	Details: "burnable, general garbage, foodwaste",
	Emoji:   "\U0001F525",
}

var plastic Garbage = Garbage{
	Type:    "plastic",
	Details: "plastic packages, general plastic garbage",
	Emoji:   "\u267B",
}

var cansAndBottles Garbage = Garbage{
	Type:    "cans and bottles",
	Details: "cans, plastic bottles, glass bottles",
	Emoji:   "\U0001F96B",
}

var CardboardAndCloth Garbage = Garbage{
	Type:    "cardboard and cloth",
	Details: "cardboard boxes and old cloth",
	Emoji:   "\U0001F4E6",
}

// collection schedule
var Schedule = map[string]Garbage{
	"Mon": plastic,
	"Tue": conbustiable,
	"Fri": cansAndBottles,
	"Sat": conbustiable,
}

func GetCollectionSchedule() string {
	return fmt.Sprintf(
		"MONDAY:\n%s\n\nTUESDAY:\n%s\n\nFRIDAY:\n%s\n%s\n\nSATURDAY:\n%s",
		Schedule["Mon"].FormatMessage(false),
		Schedule["Tue"].FormatMessage(false),
		Schedule["Fri"].FormatMessage(false),
		CardboardAndCloth.FormatMessage(false),
		Schedule["Sat"].FormatMessage(false),
	)
}

var helpReponse string = fmt.Sprintf(`
To get information from me, you need to provide weekday or garbage type. here are some examples.
garbage type %s "burnable", "general", "combustible", "cans", "bottles", "plastic"
weekday %s "Monday", "Tuesday", "Friday"
`, rightArrow, rightArrow)

// helper funcs
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
