package asciidoc

import (
  "testing"
  "time"
)

func TestGeneratorGetInfo(t *testing.T) {
  generator := NewGenerator()
  if generator == nil {
    t.Errorf("Expected generator to be non-nil")
  }
  input := `# Title

Text
`
  tm := time.Date(2019, 2, 14, 23, 0, 0, 0, time.UTC)
  generator.Parse(input, tm)
  if generator.Info.Title != "Title" {
    t.Errorf("Expected title to be `Title', got `%s'", generator.Info.Title)
  }
}

func TestGeneratorParse(t *testing.T) {
  generator := NewGenerator()
  input := `# Title

Text
`
  tm := time.Date(2019, 2, 14, 23, 0, 0, 0, time.UTC)
  generator.Parse(input, tm)
  output := generator.Output
  expected := `<?xml version="1.0" encoding="UTF-8"?>
<?asciidoc-toc?>
<?asciidoc-numbered?>
<article xmlns="http://docbook.org/ns/docbook" xmlns:xl="http://www.w3.org/1999/xlink" version="5.0" xml:lang="en">
<info>
<title>Title</title>
<date>2019-02-14</date>
</info>
<simpara>
Text
</simpara>
</article>
`
  if output != expected {
    t.Errorf("Expected output `%s', got `%s'", expected, output)
  }
}

func TestGeneratorParse2(t *testing.T) {
  generator := NewGenerator()
  input := `# Title

## Title 2

Line 1
Line 2
`
  tm := time.Date(2019, 2, 14, 23, 0, 0, 0, time.UTC)
  generator.Parse(input, tm)
  output := generator.Output
  expected := `<?xml version="1.0" encoding="UTF-8"?>
<?asciidoc-toc?>
<?asciidoc-numbered?>
<article xmlns="http://docbook.org/ns/docbook" xmlns:xl="http://www.w3.org/1999/xlink" version="5.0" xml:lang="en">
<info>
<title>Title</title>
<date>2019-02-14</date>
</info>
<section xml:id="_title_2">
<title>Title 2</title>
<simpara>
Line 1
Line 2
</simpara>
</section>
</article>
`
  if output != expected {
    t.Errorf("Expected output `%s', got `%s'", expected, output)
  }
}

func TestGeneratorWarning(t *testing.T) {
  generator := NewGenerator()
  input := `# Title

### Title 2

...
`
  tm := time.Date(2019, 2, 14, 23, 0, 0, 0, time.UTC)
  generator.Parse(input, tm)
  if !generator.Warning {
    t.Errorf("Expected warning")
  }
}

func TestGeneratorWithBold(t *testing.T) {
  generator := NewGenerator()
  input := `# Title

## Title 2

Line has something *bold
Line has something _italic
`
  tm := time.Date(2019, 2, 14, 23, 0, 0, 0, time.UTC)
  generator.Parse(input, tm)
  output := generator.Output
  expected := `<?xml version="1.0" encoding="UTF-8"?>
<?asciidoc-toc?>
<?asciidoc-numbered?>
<article xmlns="http://docbook.org/ns/docbook" xmlns:xl="http://www.w3.org/1999/xlink" version="5.0" xml:lang="en">
<info>
<title>Title</title>
<date>2019-02-14</date>
</info>
<section xml:id="_title_2">
<title>Title 2</title>
<simpara>
Line has something <emphasis role="strong">bold
Line has something <emphasis>italic
</emphasis></emphasis>
</simpara>
</section>
</article>
`
  if output != expected {
    t.Errorf("Expected output `%s', got `%s'", expected, output)
  }
}
