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

func TestHasHTMLTag(t *testing.T) {
	if !HasHTMLTag(example1b, 200) {
		t.Fail()
	}
	if !HasHTMLTag(example2b, 200) {
		t.Fail()
	}
	if HasHTMLTag(example3b, 200) {
		t.Fail()
	}
	if HasHTMLTag(example4b, 200) {
		t.Fail()
	}
	if HasHTMLTag(example5b, 200) {
		t.Fail()
	}
	if !HasHTMLTag(example6b, 200) {
		t.Fail()
	}
	if HasHTMLTag([]byte("<asdfhtml><body>hi</body></html>"), 200) {
		t.Fatal("fail: <asdfhtml><body>hi</body></html> has a html tag?")
	}
}

func TestHasScriptTag(t *testing.T) {
	if !HasScriptTag([]byte("<script>1+1</script>"), 200) {
		t.Fatal("fail: <script>1+1</script> does not have a script tag?")
	}
	if HasScriptTag([]byte("<asdfscript>1+1</script>"), 200) {
		t.Fatal("fail: <asdfscript>1+1</script> has a script tag?")
	}
}

func TestGetHTMLTag(t *testing.T) {
	if tag, err := GetHTMLTag(example1b, 200); err != nil || string(tag) != "<html>" {
		t.Fatal(string(tag))
	}
	if tag, err := GetHTMLTag(example2b, 200); err != nil || string(tag) != "<html asdf=123 qwerty=256>" {
		t.Fatal(string(tag))
	}
	if _, err := GetHTMLTag(example3b, 200); err == nil {
		t.Fail()
	}
	if _, err := GetHTMLTag(example4b, 200); err == nil {
		t.Fail()
	}
	if _, err := GetHTMLTag(example5b, 200); err == nil {
		t.Fail()
	}
	if tag, err := GetHTMLTag(example6b, 200); err != nil || string(tag) != "<html    asdf=1234>" {
		t.Fatal(string(tag))
	}
}

func TestHTMLIndex(t *testing.T) {
	if pos, err := HTMLIndex(example1b, 200); err != nil || pos != 0 {
		t.Fatal(pos)
	}
	if pos, err := HTMLIndex(example2b, 200); err != nil || pos != 0 {
		t.Fatal(pos)
	}
	if pos, err := HTMLIndex(example3b, 200); err == nil {
		t.Fatal(pos)
	}
	if pos, err := HTMLIndex(example4b, 200); err == nil {
		t.Fatal(pos)
	}
	if pos, err := HTMLIndex(example5b, 200); err == nil {
		t.Fatal(pos)
	}
	if pos, err := HTMLIndex(example6b, 200); err != nil || pos != 6 {
		t.Fatal(pos)
	}
}

func TestHasHTMLTagString(t *testing.T) {
	if !HasHTMLTagString(example1s, 200) {
		t.Fail()
	}
	if !HasHTMLTagString(example2s, 200) {
		t.Fail()
	}
	if HasHTMLTagString(example3s, 200) {
		t.Fail()
	}
	if HasHTMLTagString(example4s, 200) {
		t.Fail()
	}
	if HasHTMLTagString(example5s, 200) {
		t.Fail()
	}
	if !HasHTMLTagString(example6s, 200) {
		t.Fail()
	}
}

func TestGetHTMLTagString(t *testing.T) {
	if tag, err := GetHTMLTagString(example1s, 200); err != nil || tag != "<html>" {
		t.Fatal(tag)
	}
	if tag, err := GetHTMLTagString(example2s, 200); err != nil || tag != "<html asdf=123 qwerty=256>" {
		t.Fatal(tag)
	}
	if _, err := GetHTMLTagString(example3s, 200); err == nil {
		t.Fail()
	}
	if _, err := GetHTMLTagString(example4s, 200); err == nil {
		t.Fail()
	}
	if _, err := GetHTMLTagString(example5s, 200); err == nil {
		t.Fail()
	}
	if tag, err := GetHTMLTagString(example6s, 200); err != nil || tag != "<html    asdf=1234>" {
		t.Fatal(tag)
	}
}

func TestHTMLIndexString(t *testing.T) {
	if pos, err := HTMLIndexString(example1s, 200); err != nil || pos != 0 {
		t.Fatal(pos)
	}
	if pos, err := HTMLIndexString(example2s, 200); err != nil || pos != 0 {
		t.Fatal(pos)
	}
	if pos, err := HTMLIndexString(example3s, 200); err == nil {
		t.Fatal(pos)
	}
	if pos, err := HTMLIndexString(example4s, 200); err == nil {
		t.Fatal(pos)
	}
	if pos, err := HTMLIndexString(example5s, 200); err == nil {
		t.Fatal(pos)
	}
	if pos, err := HTMLIndexString(example6s, 200); err != nil || pos != 6 {
		t.Fatal(pos)
	}
}
