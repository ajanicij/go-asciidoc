package asciidoc

import (
  "testing"
  "time"
)

func TestParseDate(t *testing.T) {
  datestr := "1969-05-14"
  expectedDate := time.Date(1969, 5, 14, 23, 0, 0, 0, time.UTC)
  tm, err := UtilConvertDate(datestr)
  if err != nil {
    t.Errorf("Got error: %s", err)
  }
  if (tm.Year() != 1969) || (tm.Month() != 5) || (tm.Day() != 14) {
    t.Errorf("Got wrong date, expected %s and got %s", expectedDate, tm)
  }
}
