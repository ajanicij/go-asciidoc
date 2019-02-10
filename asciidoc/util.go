package asciidoc

import (
  "time"
  "strings"
  "errors"
  "strconv"
)

func UtilConvertDate(str string) (tm time.Time, err error) {
  parts := strings.Split(str, "-")
  if len(parts) != 3 {
    return time.Now(), errors.New("Invalid date format")
  }
  var year, month, day int
  year, err = strconv.Atoi(parts[0])
  if err != nil {
    return tm, err
  }
  month, err = strconv.Atoi(parts[1])
  if err != nil {
    return tm, err
  }
  day, err = strconv.Atoi(parts[2])
  if err != nil {
    return tm, err
  }
  tm = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
  return tm, nil
}
