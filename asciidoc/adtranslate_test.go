package asciidoc

import "testing"

func TestDummy(t *testing.T) {
  x1 := 1
  x2 := 1
  if x1 != x2 {
    t.Errorf("Expected x1=x2, but x1=%d, x2=%d", x1, x2)
  }
}

func TestStateReset(t *testing.T) {
  state := NewState()
  state.Event_heading = true
  state.Event_paragraph = true
  state.Reset()
  if state.Event_heading {
    t.Errorf("Expected Event_heading=false, got true")
  }
  if state.Event_paragraph {
    t.Errorf("Expected Event_paragraph=false, got true")
  }
}

func TestHeadingState(t *testing.T) {
  state := &HeadingState{state: 0}
  state.push("= heading 0")
  if state.state != 1 {
    t.Errorf("Expected state=1, got %d", state.state)
  }
  state.push("")
  if state.state != 0 {
    t.Errorf("Expected state=0, got %d", state.state)
  }
  state.push("text")
  if state.state != 2 {
    t.Errorf("Expected state=2, got %d", state.state)
  }
  state.push("text 2")
  if state.state != 3 {
    t.Errorf("Expected state=3, got %d", state.state)
  }
  state.push("")
  if state.state != 4 {
    t.Errorf("Expected state=4, got %d", state.state)
  }
  state.push("Heading level 1")
  if state.state != 2 {
    t.Errorf("Expected state=2, got %d", state.state)
  }
  state.push("---------------")
  if state.state != 1 {
    t.Errorf("Expected state=1, got %d", state.state)
  }
}

func TestGetHeading(t *testing.T) {
  state := NewState()
  state.Push("= Heading")
  if !state.Event_heading {
    t.Errorf("Expected Event_heading to be true")
  }
  if state.State.Cell_heading != "Heading" {
    t.Errorf("Expected Cell_heading to be `Heading', got %s", state.State.Cell_heading)
  }
  state.Push("")
  state.Push("== Heading level 1")
  if !state.Event_heading {
    t.Errorf("Expected Event_heading to be true")
  }
  state.Push("")
  state.Push("=== Heading level 2")
  if !state.Event_heading {
    t.Errorf("Expected Event_heading to be true")
  }
  if state.State.Cell_headingLevel != 2 {
    t.Errorf("Expected Cell_headingLevel to be 2, got %d", state.State.Cell_headingLevel)
  }
  state.Push("")
  state.Push("Heading level 3")
  state.Push("^^^^^^^^^^^^^^^")
  if !state.Event_heading {
    t.Errorf("Expected Event_heading to be true")
  }
  if state.State.Cell_headingLevel != 3 {
    t.Errorf("Expected Cell_headingLevel to be 3, got %d", state.State.Cell_headingLevel)
  }
  if state.State.Cell_heading != "Heading level 3" {
    t.Errorf("Expected Cell_heading to be `Heading level 3', got %s", state.State.Cell_heading)
  }

  // Check Event_heading is false when it should be.
  state.Push("xyz")
  if state.Event_heading {
    t.Errorf("Expected Event_heading to be false, got true")
  }

  state.Push("")
  state.Push("## Heading level 1")
  if !state.Event_heading {
    t.Errorf("Expected Event_heading to be true")
  }
  if state.State.Cell_heading != "Heading level 1" {
    t.Errorf("Expected Cell_heading to be `Heading level 1', got %s", state.State.Cell_heading)
  }
}

func equalParagraphs(p1, p2 []string) bool {
  if p1 == nil {
    if p2 == nil {
      return true
    }
    return false
  }
  if p2 == nil {
    return false
  }
  if len(p1) != len(p2) {
    return false
  }
  for i, el := range(p1) {
    if el != p2[i] {
      return false
    }
  }
  return true
}

func TestGetParagraph(t *testing.T) {
  state := NewState()
  state.Push("= Heading")
  state.Push("")
  state.Push("line 1")
  state.Push("line 2")
  state.Push("")
  if !state.Event_paragraph {
    t.Errorf("Expected Event_paragraph to be true, got false")
  }
  if !equalParagraphs(state.State.Cell_paragraph, []string{"line 1", "line 2"}) {
    t.Errorf("Expected Cell_paragraph to be line 1/line 2, got %s", state.State.Cell_paragraph)
  }
  // TODO: Check for one-line paragraph.
  state.Push("line 1")
  state.Push("--")
  state.Push("line 3")
  state.Push("")
  if !state.Event_paragraph {
    t.Errorf("Expected Event_paragraph to be true, got false")
  }
  if !equalParagraphs(state.State.Cell_paragraph, []string{"line 1", "--", "line 3"}) {
    t.Errorf("Expected Cell_paragraph to be line 1/--/line 2, got %s", state.State.Cell_paragraph)
  }
}

func TestStack(t *testing.T) {
  stack := NewStack()
  if stack == nil {
    t.Errorf("Expected NewStack to create a stack")
  }
  stack.Push(1001)
  if stack.Empty() {
    t.Errorf("stack should not be empty")
  }
  // TODO: check err returned by Top() and Pop().
  x, _ := stack.Top()
  if x != 1001 {
    t.Errorf("Expected top of stack to be 1001, got %d", x)
  }
  x, _ = stack.Pop()
  if x != 1001 {
    t.Errorf("Expected top of stack to be 1001, got %d", x)
  }
  if !stack.Empty() {
    t.Errorf("Expected stack to be empty")
  }
}
