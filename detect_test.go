package huldra

import (
	"testing"
)

const (
	example1s = "<html><body>hi</body></html>"
	example2s = "<html asdf=123 qwerty=256><body>hi</body></html>"
	example3s = "asdf"
	example4s = "<html"
	example5s = "                           \n\n\n\n\n<html"
	example6s = "\n\n\n   <html    asdf=1234>   \n\n\n<body>hi</body></html><html></html>"
)

var (
	example1b = []byte(example1s)
	example2b = []byte(example2s)
	example3b = []byte(example3s)
	example4b = []byte(example4s)
	example5b = []byte(example5s)
	example6b = []byte(example6s)
)

func TestIsHTML(t *testing.T) {
	if !IsHTML(example1b) {
		t.Fail()
	}
	if !IsHTML(example2b) {
		t.Fail()
	}
	if IsHTML(example3b) {
		t.Fail()
	}
	if !IsHTML(example4b) {
		t.Fail()
	}
	if !IsHTML(example5b) {
		t.Fail()
	}
	if !IsHTML(example6b) {
		t.Fail()
	}
}

func TestGetHTMLTag(t *testing.T) {
	if tag, err := GetHTMLTag(example1b); err != nil || string(tag) != "<html>" {
		t.Fatal(string(tag))
	}
	if tag, err := GetHTMLTag(example2b); err != nil || string(tag) != "<html asdf=123 qwerty=256>" {
		t.Fatal(string(tag))
	}
	if _, err := GetHTMLTag(example3b); err == nil {
		t.Fail()
	}
	if tag, err := GetHTMLTag(example4b); err != nil || string(tag) != "<html" {
		t.Fatal(string(tag))
	}
	if tag, err := GetHTMLTag(example5b); err != nil || string(tag) != "<html" {
		t.Fatal(string(tag))
	}
	if tag, err := GetHTMLTag(example6b); err != nil || string(tag) != "<html    asdf=1234>" {
		t.Fatal(string(tag))
	}
}

func TestIsHTMLString(t *testing.T) {
	if !IsHTMLString(example1s) {
		t.Fail()
	}
	if !IsHTMLString(example2s) {
		t.Fail()
	}
	if IsHTMLString(example3s) {
		t.Fail()
	}
	if !IsHTMLString(example4s) {
		t.Fail()
	}
	if !IsHTMLString(example5s) {
		t.Fail()
	}
	if !IsHTMLString(example6s) {
		t.Fail()
	}
}

func TestGetHTMLTagString(t *testing.T) {
	if tag, err := GetHTMLTagString(example1s); err != nil || tag != "<html>" {
		t.Fatal(tag)
	}
	if tag, err := GetHTMLTagString(example2s); err != nil || tag != "<html asdf=123 qwerty=256>" {
		t.Fatal(tag)
	}
	if _, err := GetHTMLTagString(example3s); err == nil {
		t.Fail()
	}
	if tag, err := GetHTMLTagString(example4s); err != nil || tag != "<html" {
		t.Fatal(tag)
	}
	if tag, err := GetHTMLTagString(example5s); err != nil || tag != "<html" {
		t.Fatal(tag)
	}
	if tag, err := GetHTMLTagString(example6s); err != nil || tag != "<html    asdf=1234>" {
		t.Fatal(tag)
	}
}
