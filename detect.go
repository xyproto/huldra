package huldra

import (
	"errors"
	"strings"
)

var errNoHTML = errors.New("could not find a <html> tag")

func IsHTML(data []byte) bool {
	var foundCounter uint8
	for i, r := range data {
		switch r {
		case '<':
			foundCounter = 1
		case 'h', 'H':
			if foundCounter == 1 {
				foundCounter++
			}
		case 't', 'T':
			if foundCounter == 2 {
				foundCounter++
			}
		case 'm', 'M':
			if foundCounter == 3 {
				foundCounter++
			}
		case 'l', 'L':
			if foundCounter == 4 {
				foundCounter++
			}
		case ' ', '>':
			if foundCounter == 5 {
				return true // found "<html " or "<html>"
			}
		}
		if i > 200 {
			return false // no <html> tag for the first 200 bytes
		}
	}
	return false
}

func HTMLIndex(data []byte) int {
	var (
		foundCounter uint8
		pos          int
	)
	for i, r := range data {
		switch r {
		case '<':
			pos = i
			foundCounter = 1
		case 'h', 'H':
			if foundCounter == 1 {
				foundCounter++
			}
		case 't', 'T':
			if foundCounter == 2 {
				foundCounter++
			}
		case 'm', 'M':
			if foundCounter == 3 {
				foundCounter++
			}
		case 'l', 'L':
			if foundCounter == 4 {
				foundCounter++
			}
		case ' ', '>':
			if foundCounter == 5 {
				// found "<html " or "<html>"
				return pos
			}
		}
		if pos > 200 {
			return -1 // no <html> tag for the first 200 bytes
		}
	}
	return -1
}

func IsHTMLString(s string) bool {
	var foundCounter uint8
	for i, r := range s {
		switch r {
		case '<':
			foundCounter = 1
		case 'h', 'H':
			if foundCounter == 1 {
				foundCounter++
			}
		case 't', 'T':
			if foundCounter == 2 {
				foundCounter++
			}
		case 'm', 'M':
			if foundCounter == 3 {
				foundCounter++
			}
		case 'l', 'L':
			if foundCounter == 4 {
				foundCounter++
			}
		case ' ', '>':
			if foundCounter == 5 {
				// found "<html " or "<html>"
				return true
			}
		}
		if i > 200 {
			return false // no <html> tag for the first 200 runes
		}
	}
	return false
}

func HTMLIndexString(s string) int {
	var (
		foundCounter uint8
		pos          int
	)
	for i, r := range s {
		switch r {
		case '<':
			pos = i
			foundCounter = 1
		case 'h', 'H':
			if foundCounter == 1 {
				foundCounter++
			}
		case 't', 'T':
			if foundCounter == 2 {
				foundCounter++
			}
		case 'm', 'M':
			if foundCounter == 3 {
				foundCounter++
			}
		case 'l', 'L':
			if foundCounter == 4 {
				foundCounter++
			}
		case ' ', '>':
			if foundCounter == 5 {
				// found "<html " or "<html>"
				return pos
			}
		}
		if pos > 200 {
			return -1 // no <html> tag for the first 200 runes
		}
	}
	return -1
}

func GetHTMLTag(data []byte) ([]byte, error) {
	pos := HTMLIndex(data)
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
	pos := HTMLIndexString(s)
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
