package weekdays

import "testing"

func TestMergeWeekdays(t *testing.T) {
	w := Weekdays(Sunday | Friday)

	w2 := Weekdays(Wednesday | Thursday)

	w = w.Set(w2)
	w = w.Set(Monday)

	if w.String() != "sun, mon, wed, thu, fri" {
		t.Errorf("Unexpected string: %s", w.String())
	}
}
