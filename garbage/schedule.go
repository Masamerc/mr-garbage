package garbage

import (
	"fmt"
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
