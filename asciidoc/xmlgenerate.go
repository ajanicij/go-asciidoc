package asciidoc

import (
  "fmt"
  "time"
  "unicode"
  "strings"
)

// GenerateDateString takes a Time instance and generates date string like
// this: "2019-02-04". To generate a date string for today's date, call
// GenerateDateString(time.Now())
func GenerateDateString(t time.Time) string {
  return fmt.Sprintf("%04d-%02d-%02d", t.Year(), t.Month(), t.Day())
}

func GenerateId(str string) string {
  filtered := ""
  for _, ch := range(str) {
    if unicode.IsDigit(ch) || unicode.IsLetter(ch) {
      filtered = filtered + string(unicode.ToLower(ch))
    } else if unicode.IsSpace(ch) {
      filtered = filtered + " "
    }
  }
  words := strings.Split(filtered, " ")
  result := ""
  for _, word := range words {
    result = result + "_" + word
  }
  return result
}

type StringSet map[string]bool

func (s StringSet) IsEmpty() bool {
  return (len(s) == 0)
}

func (s StringSet) Len() int {
  return len(s)
}

func NewStringSet() *StringSet {
  set := make(map[string]bool)
  return (*StringSet)(&set)
}

func (s StringSet) Contains(str string) bool {
  set := (map[string]bool)(s)
  _, ok := set[str]
  return ok
}

func (s *StringSet) Add(str string) {
  set := (map[string]bool)(*s)
  set[str] = true
}

func (s *StringSet) GenerateUniqueId(str string) string {
  str = GenerateId(str)

  if !s.Contains(str) {
    s.Add(str)
    return str
  }

  for i := 2; ; i++ {
    trystr := fmt.Sprintf("%s_%d", str, i)
    if !s.Contains(trystr) {
      s.Add(trystr)
      return trystr
    }
  }
}
