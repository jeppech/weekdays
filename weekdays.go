package weekdays

import "strings"

type Weekday uint8

// Weekdays represents a bitmask of weekdays, where bit 0 is Sunday
type Weekdays Weekday

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (w Weekdays) Set(day Weekday) Weekdays {
	w |= 1 << day
	return w
}

func (w Weekdays) Unset(day Weekday) Weekdays {
	w &= ^(1 << day)
	return w
}

func (w Weekdays) Merge(days Weekdays) Weekdays {
	w |= days
	return w
}

func (w Weekdays) IsSet(day Weekday) bool {
	return w&(1<<day) != 0
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
	return strings.Join(ret, ", ")
}
