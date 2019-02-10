package asciidoc

import (
  "bufio"
  "strings"
)

type Info struct {
  Title string
}

func GetInfo(s string) *Info {
  info := &Info{}
  reader := strings.NewReader(s)
  scanner := bufio.NewScanner(reader)
  state := NewState()
  for scanner.Scan() {
    line := scanner.Text()
    state.Push(line)
    if state.Event_heading {
      info.Title = state.State.Cell_heading
      break
    }
  }
  return info
}
