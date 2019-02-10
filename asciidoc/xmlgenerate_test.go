package asciidoc

import (
  "testing"
  "time"
)

func TestGenerateDateString(t *testing.T) {
  tobj := time.Date(1969, 5, 14, 0, 0, 0, 0, time.UTC)
  date_string := GenerateDateString(tobj)
  if date_string != "1969-05-14" {
    t.Errorf("Expected date string=`1969-05-14', got %s", date_string)
  }
}

func TestGenerateId(t *testing.T) {
  str := "two words"
  id := GenerateId(str)
  if id != "_two_words" {
    t.Errorf("Expected id=`_two_words', got `%s'", id)
  }

  // TODO: check GenerateId from "new \" title" (string that contains characters that
  //       are neither digit nor letter)
}

func TestStringSet(t *testing.T) {
  set := NewStringSet()
  if !set.IsEmpty() {
    t.Errorf("Expected empty set")
  }
}

func TestStringSetContains(t *testing.T) {
  set := NewStringSet()
  if set.Contains("abc") {
    t.Errorf("Expected that set does not contain `abc'")
  }
}

func TestStringSetAdd(t *testing.T) {
  set := NewStringSet()
  set.Add("abc")
  if !set.Contains("abc") {
    t.Errorf("Expected that set contains `abc'")
  }
}

func TestGenerateUniqueId(t *testing.T) {
  set := NewStringSet()
  set.Add("_abc")
  set.Add("_abc_2")
  str := "abc"
  id := set.GenerateUniqueId(str)
  if id != "_abc_3" {
    t.Errorf("Expected id=`_abc_3', got %s", id)
  }

  str = "new title"
  id = set.GenerateUniqueId(str)
  if id != "_new_title" {
    t.Errorf("Expected id=`_new_title', got %s", id)
  }

  id = set.GenerateUniqueId(str)
  if id != "_new_title_2" {
    t.Errorf("Expected id=`_new_title_2', got %s", id)
  }
}
