package asciidoc

import "testing"

func TestConvertBold(t *testing.T) {
  converter := NewConverter()
  instr := "a *bold* text"
  outstr := converter.Convert(instr)
  expected := "a <emphasis role=\"strong\">bold</emphasis> text"
  if outstr != expected {
    t.Errorf("Expected `%s', got `%s'", expected, outstr)
  }
}

func TestConvertItalic(t *testing.T) {
  converter := NewConverter()
  instr := "a _italic_ text"
  outstr := converter.Convert(instr)
  expected := "a <emphasis>italic</emphasis> text"
  if outstr != expected {
    t.Errorf("Expected `%s', got `%s'", expected, outstr)
  }
}

func TestConvertLines(t *testing.T) {
  converter := NewConverter()
  input := []string{"a *bold* line", "an _italic_ line", "and now a *bold", "trick* coming"}
  expected := []string{
    "a <emphasis role=\"strong\">bold</emphasis> line",
    "an <emphasis>italic</emphasis> line",
    "and now a <emphasis role=\"strong\">bold",
    "trick</emphasis> coming"}
  output := converter.ConvertLines(input)
  if !equalParagraphs(expected, output) {
    t.Errorf("Expected output to be %s, got %s", expected, output)
  }
}

func TestConvertEnd(t *testing.T) {
  converter := NewConverter()
  input := []string{
    "line with *bold and _italic"}
  expected := []string{
    "line with <emphasis role=\"strong\">bold and <emphasis>italic",
    "</emphasis></emphasis>"}
  output := converter.ConvertLines(input)
  if !equalParagraphs(expected, output) {
    t.Errorf("Expected output to be %s, got %s", expected, output)
  }
  if !converter.Warning {
    t.Errorf("Expected warning")
  }
}
