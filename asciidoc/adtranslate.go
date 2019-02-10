package asciidoc

import (
  "fmt"
  "strings"
)

type State struct {
  Input string
  Event_heading bool
  Event_paragraph bool
  State *HeadingState
}

type HeadingState struct {
  state int
  firstLine string
  Cell_heading string
  Cell_headingLevel int
  Cell_paragraph []string
}

func newHeadingState() *HeadingState {
  state := &HeadingState{}
  state.reset()
  return state
}

func (state *HeadingState) reset() {
  state.state = 0
  state.firstLine = ""
  state.Cell_heading = ""
  state.Cell_headingLevel = 0
  state.Cell_paragraph = nil
}

func NewState() *State {
  state := &State{}
  state.Reset()
  return state
}

func (state *State) Reset() {
  state.Input = ""
  state.Event_heading = false
  state.Event_paragraph = false
  state.State = newHeadingState()
}

func Dummy() {
  fmt.Println("in adtranslate#Dummy")
}

// Heading-and-paragraph detection state machine
// Input is line of text. Possible checks:
// - line is empty
// - line is nonempty
// - line is a header line ("= ...")
// - line is an underline ("--...-")
// States are between lines:
//   (state 0)
// = Heading level 0
//   (state 1)
// <empty line>
//   (state 0)
// Text
//   (state 2)
// Text
//   (state 3)
// <empty line>
//   (state 4)
// Heading level 1
//   (state 2)
// ---------------
//   (state 1)
// <empty line>
//   (state 0)
// Heading level 1
//   (state 2)
// ---------------
//   (state 1)
// Text
//   (state 3)

func isHeadingLine(line string) (res bool, level int, heading string) {
  // TODO: real implementation
  if strings.HasPrefix(line, "= ") {
    return true, 0, strings.TrimPrefix(line, "= ")
  }
  if strings.HasPrefix(line, "# ") {
    return true, 0, strings.TrimPrefix(line, "# ")
  }
  if strings.HasPrefix(line, "== ") {
    return true, 1, strings.TrimPrefix(line, "== ")
  }
  if strings.HasPrefix(line, "## ") {
    return true, 1, strings.TrimPrefix(line, "## ")
  }
  if strings.HasPrefix(line, "=== ") {
    return true, 2, strings.TrimPrefix(line, "=== ")
  }
  if strings.HasPrefix(line, "### ") {
      return true, 2, strings.TrimPrefix(line, "### ")
  }
  if strings.HasPrefix(line, "==== ") {
    return true, 3, strings.TrimPrefix(line, "==== ")
  }
  if strings.HasPrefix(line, "#### ") {
    return true, 3, strings.TrimPrefix(line, "#### ")
  }
  if strings.HasPrefix(line, "===== ") {
    return true, 4, strings.TrimPrefix(line, "===== ")
  }
  if strings.HasPrefix(line, "##### ") {
    return true, 4, strings.TrimPrefix(line, "##### ")
  }

  return false, 0, ""
}

// isAllChars checks if an input string consists of the same character.
// Returns true if line is non-empty and it only contains ch.
func isAllChars(line string, ch rune) bool {
  if len(line) == 0 {
    return false
  }
  for _, c := range(line) {
    if c != ch {
      return false
    }
  }
  return true
}

func isHeadingUnderline(line string) (res bool, level int) {
  if isAllChars(line, '-') {
    return true, 1
  }
  if isAllChars(line, '~') {
    return true, 2
  }
  if isAllChars(line, '^') {
    return true, 3
  }
  if isAllChars(line, '^') {
    return true, 4
  }
  return false, 0
}

func abs(n int) int {
  if n < 0 {
    return -n
  }
  return n
}

// push pushes one input line to state machine that tracks the state
// of headings/paragraphs.
// TODO: all state transitions
func (state *HeadingState) push(line string) {
  var isHeading bool
  var level int
  var text string

  switch state.state {
  case 0, 4:
    isHeading, level, text = isHeadingLine(line)
    if isHeading {
      state.state = 1
      state.Cell_heading = text
      state.Cell_headingLevel = level
    } else if !isHeading {
      if line != "" {
        state.state = 2
        state.firstLine = line
        state.Cell_paragraph = []string{line}
      } else {
        state.state = 0
      }
    }
  case 1:
    if line == "" {
      state.state = 0
    } else {
      state.state = 3
    }
  case 2:
    if line != "" {
      isHeading, level = isHeadingUnderline(line)
      if isHeading && (abs(len(state.firstLine) - len(line)) <= 1) {
        state.state = 1
        state.Cell_heading = state.firstLine
        state.Cell_headingLevel = level
      } else {
        state.state = 3
        state.Cell_paragraph = append(state.Cell_paragraph, line)
      }
    }
    if line == "" {
      state.state = 4
    }
  case 3:
    if line == "" {
      state.state = 4
    } else {
      state.Cell_paragraph = append(state.Cell_paragraph, line)
    }
  }
}

func (state *State) Push(line string) {
  state.Input = line
  state.State.push(line)
  state.Event_heading = (state.State.state == 1)
  state.Event_paragraph = (state.State.state == 4)
}
