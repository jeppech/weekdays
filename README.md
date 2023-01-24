# weekdays
A package for tinkering with weekdays.

```go
package weekdays

func test() {
  w := Weekdays(0)

  w.Set(Sunday)
  w.Set(Friday)

  if w.IsSet(Sunday) && w.IsSet(Friday) {
    println(w) // "fri,sun"
  }

  w2 := Weekdays(0)
  w2.Set(Wednesday)
  w2.Set(Thursday)

  if w2.IsSet(Wednesday) && w2.IsSet(Thursday) {
    w2.Merge(w)
    println(w2) // "fri,sun,thu,wed"
  }
}
```
