package asciidoc

import (
  "time"
  "strings"
  "bufio"
  "fmt"
)

type Generator struct {
  Info *Info
  Output string
  stack *Stack
  set *StringSet
  Warning bool
  WarningList []string
  CurrentLine int
}

func NewGenerator() *Generator {
  generator := &Generator{stack: NewStack(), set: NewStringSet()}
  return generator
}

func (generator *Generator) Parse(input string, tm time.Time) {
  input = input + "\n" // just in case, append a newline
  generator.Info = GetInfo(input)
  var buf strings.Builder
  b := bufio.NewWriter(&buf)
  generator.WriteHead(b, tm)
  reader := strings.NewReader(input)
  scanner := bufio.NewScanner(reader)
  state := NewState()
  for scanner.Scan() {
    line := scanner.Text()
    state.Push(line)
    generator.CurrentLine++
    if state.Event_heading {
      generator.OnEventHeading(b, state.State.Cell_headingLevel, state.State.Cell_heading)
    } else if state.Event_paragraph {
      generator.OnEventParagraph(b, state.State.Cell_paragraph)
    }
  }
  for !generator.stack.Empty() {
    generator.stack.Pop()
    fmt.Fprintf(b, "</section>\n")
  }
  fmt.Fprintf(b, "</article>\n")
  b.Flush()
  generator.Output = buf.String()
}

func (generator *Generator) WriteHead(writer *bufio.Writer, tm time.Time) {
  fmt.Fprintf(writer, `<?xml version="1.0" encoding="UTF-8"?>
<?asciidoc-toc?>
<?asciidoc-numbered?>
<article xmlns="http://docbook.org/ns/docbook" xmlns:xl="http://www.w3.org/1999/xlink" version="5.0" xml:lang="en">
<info>
<title>%s</title>
<date>%s</date>
</info>
`, generator.Info.Title, GenerateDateString(tm))
}

func (generator *Generator) OnEventHeading(writer *bufio.Writer, level int, heading string) {
  var current int
  if level == 0 {
    return
  }
  for !generator.stack.Empty() {
    current, _ = generator.stack.Top()
    if current < level {
      break
    }
    generator.stack.Pop()
    fmt.Fprintf(writer, "</section>\n")
  }
  if level > current + 1 {
    generator.Warning = true
    generator.AddWarning(fmt.Sprintf("section title out of sequence: expected level %d, got level %d", current + 1, level))
  }
  fmt.Fprintf(writer, "<section xml:id=\"%s\">\n", generator.set.GenerateUniqueId(heading))
  generator.stack.Push(level)
  fmt.Fprintf(writer, "<title>%s</title>\n", heading)
}

func (generator *Generator) OnEventParagraph(writer *bufio.Writer, paragraph []string) {
  converter := NewConverter()
  lines := converter.ConvertLines(paragraph)
  fmt.Fprintf(writer, "<simpara>\n")
  for _, paraline := range lines {
    fmt.Fprintf(writer, "%s\n", paraline)
  }
  fmt.Fprintf(writer, "</simpara>\n")
  if converter.Warning {
    for _, message := range converter.WarningList {
      generator.AddWarning(message)
    }
  }
}

func (generator *Generator) AddWarning(message string) {
  message = fmt.Sprintf("line %d: %s", generator.CurrentLine, message)
  generator.WarningList = append(generator.WarningList, message)
}
