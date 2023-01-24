package weekdays

import "strings"

// Weekdays represents a bitmask of weekdays, where bit 0 is Sunday
type Weekdays uint8

const (
	Sunday Weekdays = 1 << iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// Set sets the given weekday
func (w Weekdays) Set(day Weekdays) Weekdays {
	return w | day
}

func (w Weekdays) Unset(day Weekdays) Weekdays {
	return w &^ day
}

func (w Weekdays) IsSet(day Weekdays) bool {
	return w&day != 0
}

func (w Weekdays) String() string {
	var ret []string

	if w.IsSet(Sunday) {
		ret = append(ret, "sun")
	}
	if w.IsSet(Monday) {
		ret = append(ret, "mon")
	}
	if w.IsSet(Tuesday) {
		ret = append(ret, "tue")
	}
	if w.IsSet(Wednesday) {
		ret = append(ret, "wed")
	}
	if w.IsSet(Thursday) {
		ret = append(ret, "thu")
	}
	if w.IsSet(Friday) {
		ret = append(ret, "fri")
	}
	if w.IsSet(Saturday) {
		ret = append(ret, "sat")
	}
	return strings.Join(ret, ",")
}
