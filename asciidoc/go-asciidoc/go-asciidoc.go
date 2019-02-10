package main

import (
  "fmt"
  "log"
  "os"
  "github.com/ajanicij/go-asciidoc/asciidoc"
  "io/ioutil"
  "time"
  "path/filepath"
  "strings"
  "flag"
)

func usage() {
  fmt.Printf(`usage: go-asciidoc <input-file> [ --date <date> ]
Date is in format: yyyy-mm-dd (e.g. 2019-05-14)
`)
  flag.PrintDefaults()
  os.Exit(0)
}

func generateOutputFilename(in string) string {
  ext := filepath.Ext(in)
  result := in
  if ext != "" {
    result = strings.TrimSuffix(in, ext)
  }
  result = result + ".xml"
  return result
}

var setDate string

func init() {
  flag.StringVar(&setDate, "date", "", "hard-coded date in format yyyy-mm-dd")
}

func main() {
  flag.Usage = usage
  flag.Parse()
  var tm time.Time
  var err error
  if setDate != "" {
    log.Printf("Date hard-coded to %s\n", setDate)
    tm, err = test10.UtilConvertDate(setDate)
    if err != nil {
      log.Fatal(err)
    }
  } else {
    tm = time.Now()
  }

  if len(flag.Args()) < 1 {
    flag.Usage()
  }

  infile := flag.Arg(0)

  outfilename := generateOutputFilename(infile)
  outfile, err := os.Create(outfilename)
  if err != nil {
    fmt.Errorf("%s", err)
  }
  defer outfile.Close()

  filedata, err := ioutil.ReadFile(infile)
  if err != nil {
    log.Fatal(err)
  }
  filestr := string(filedata)
  generator := test10.NewGenerator()
  generator.Parse(filestr, tm)
  if generator.Warning {
    for _, message := range generator.WarningList {
      log.Println(message)
    }
  }
  fmt.Fprintf(outfile, "%s\n", generator.Output)
}
