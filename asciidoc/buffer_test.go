package asciidoc

import (
  "testing"
  "bufio"
  "strings"
  "fmt"
)

func TestBufferWrite(t *testing.T) {
  var buf strings.Builder
  s := buf.String()
  if s != "" {
    t.Errorf("Expected buffer to be empty, but it is %s", s)
  }
  b := bufio.NewWriter(&buf)
  fmt.Fprintf(b, "x = %d", 10)
  b.Flush()
  s = buf.String()
  if s != "x = 10" {
    t.Errorf("Expected buffer to contain `x = 10', it is `%s'", s)
  }
}
