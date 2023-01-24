# weekdays
A package for tinkering with weekdays.  

The underlying type of `Weekdays` is a regular `uint8`.

```go
package weekdays

import (
  "github.com/jeppech/weekdays"
)

func test() {
  w := weekdays.Weekdays(0)

  w.Set(weekdays.Sunday)
  w.Set(weekdays.Friday)

  if w.IsSet(weekdays.Sunday) && w.IsSet(weekdays.Friday) {
    fmt.Println(w) // "fri,sun"
  }

  w2 := weekdays.Weekdays(0)
  w2.Set(weekdays.Wednesday)
  w2.Set(weekdays.Thursday)

  if w2.IsSet(weekdays.Wednesday) && w2.IsSet(weekdays.Thursday) {
    w2.Merge(w)
    fmt.Println(w2) // "fri,sun,thu,wed"
  }
}
```
