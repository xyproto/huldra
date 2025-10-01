package huldra

import (
	"bytes"
	"errors"
	"strings"
)

const htmlTagStartString = "<html"

var (
	errNoHTML         = errors.New("could not find a <html> tag")
	htmlTagStartBytes = []byte(htmlTagStartString)
)

func IsHTML(data []byte) bool {
	// TODO: Make this more robust, and with better performance (return early)
	return bytes.Contains(data, htmlTagStartBytes)
}

func IsHTMLString(s string) bool {
	// TODO: Make this more robust, and with better performance (return early)
	return strings.Contains(s, htmlTagStartString)
}

func GetHTMLTag(data []byte) ([]byte, error) {
	pos := bytes.Index(data, htmlTagStartBytes)
	if pos == -1 {
		return nil, errNoHTML
	}
	var collected []byte
	collected = append(collected, '<')
	for _, r := range data[pos+1:] {
		if r == '>' { // encountered the end of the tag
			collected = append(collected, r)
			break
		} else if r == '<' { // encountered another tag
			break
		}
		// continue collecting the contents of the <html ...> tag
		collected = append(collected, r)
	}
	return collected, nil
}

func GetHTMLTagString(s string) (string, error) {
	pos := strings.Index(s, htmlTagStartString)
	if pos == -1 {
		return "", errNoHTML
	}
	var sb strings.Builder
	sb.WriteRune('<')
	for _, r := range s[pos+1:] {
		if r == '>' { // encountered the end of the tag
			sb.WriteRune(r)
			break
		} else if r == '<' { // encountered another tag
			break
		}
		// continue collecting the contents of the <html ...> tag
		sb.WriteRune(r)
	}
	return sb.String(), nil
}
