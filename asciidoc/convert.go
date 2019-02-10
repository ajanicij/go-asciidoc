package asciidoc

import (
  "strings"
  "bufio"
  "fmt"
)

type Converter struct {
  buf strings.Builder
  writer *bufio.Writer
  bold bool
  italic bool
  Warning bool
  WarningList []string
}

func NewConverter() *Converter {
  converter := &Converter{}
  converter.Reset()
  return converter
}

func (converter *Converter) Reset() {
  converter.bold = false
  converter.italic = false
  converter.ClearOutput()
  converter.Warning = false
  converter.WarningList = nil
}

func (converter *Converter) ClearOutput() {
  converter.buf.Reset()
  converter.writer = bufio.NewWriter(&converter.buf)
}

func (converter *Converter) Convert(str string) string {
  for _, ch := range str {
    converter.Push(ch)
  }
  output := converter.GetOutput()
  converter.ClearOutput()
  return output
}

func (converter *Converter) ConvertLines(lines []string) []string {
  converter.Reset()
  var output []string
  for _, line := range lines {
    outline := converter.Convert(line)
    output = append(output, outline)
  }
  line := converter.ConvertEnd()
  if line != "" {
    output = append(output, line)
  }
  return output
}

func (converter *Converter) ConvertEnd() string {
  var output string
  if converter.bold {
    output = output + "</emphasis>"
    converter.bold = false
    converter.AddWarning("Didn't see end of bold")
  }
  if converter.italic {
    output = output + "</emphasis>"
    converter.italic = false
    converter.AddWarning("Didn't see end of italic")
  }
  return output
}

func (converter *Converter) Push(ch rune) {
  if ch == '*' {
    if !converter.bold {
      converter.Emit("<emphasis role=\"strong\">")
    } else {
      converter.Emit("</emphasis>")
    }
    converter.bold = !converter.bold
  } else if ch == '_' {
    if !converter.italic {
      converter.Emit("<emphasis>")
    } else {
      converter.Emit("</emphasis>")
    }
    converter.italic = !converter.italic
  } else {
    converter.EmitChar(ch)
  }
}

func (converter *Converter) Emit(str string) {
  fmt.Fprintf(converter.writer, "%s", str)
}

func (converter *Converter) EmitChar(ch rune) {
  converter.Emit(string(ch))
}

func (converter *Converter) GetOutput() string {
  converter.writer.Flush()
  return converter.buf.String()
}

func (converter *Converter) AddWarning(message string) {
  converter.WarningList = append(converter.WarningList, message)
  converter.Warning = true
}
