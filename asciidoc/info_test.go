package asciidoc

import (
  "testing"
)

// Using Info:
// info := GetInfo(s string)
// Get info.Title - document title (level 0 heading)

func TestGetInfo(t *testing.T) {
  info := GetInfo(`xyz

# title

text
`)
  if info.Title != "title" {
    t.Errorf("Expected title to be `title', got `%s'", info.Title)
  }
}
