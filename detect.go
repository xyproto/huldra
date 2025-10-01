package huldra

import (
	"errors"
	"strings"
)

var errNoHTML = errors.New("could not find a <html> tag")

func HasHTMLTag(data []byte) bool {
	var (
		foundCounter uint8
		i            uint64
	)
	for _, r := range data {
		switch foundCounter {
		case 1:
			if r == 'h' || r == 'H' {
				foundCounter++
			}
		case 2:
			if r == 't' || r == 'T' {
				foundCounter++
			}
		case 3:
			if r == 'm' || r == 'M' {
				foundCounter++
			}
		case 4:
			if r == 'l' || r == 'L' {
				foundCounter++
			}
		case 5:
			if r == '>' || r == ' ' {
				// found "<html " or "<html>"
				return true
			}
		default:
			if r == '<' {
				foundCounter = 1
			} else {
				foundCounter = 0
			}
		}
		if i > 200 {
			return false // no <html> tag for the first 200 bytes
		}
		i++
	}
	return false
}

func HasScriptTag(data []byte) bool {
	var (
		foundCounter uint8
		i            uint64
	)
	for _, r := range data {
		switch foundCounter {
		case 1:
			if r == 's' || r == 'S' {
				foundCounter++
			}
		case 2:
			if r == 'c' || r == 'C' {
				foundCounter++
			}
		case 3:
			if r == 'r' || r == 'R' {
				foundCounter++
			}
		case 4:
			if r == 'i' || r == 'I' {
				foundCounter++
			}
		case 5:
			if r == 'p' || r == 'P' {
				foundCounter++
			}
		case 6:
			if r == 't' || r == 'T' {
				foundCounter++
			}
		case 7:
			if r == '>' || r == ' ' {
				// found "<script " or "<script>"
				return true
			}
		default:
			if r == '<' {
				foundCounter = 1
			} else {
				foundCounter = 0
			}
		}
		i++
	}
	return false
}

func HTMLIndex(data []byte) (uint64, error) {
	var (
		foundCounter uint8
		i            uint64
		pos          uint64
	)
	for _, r := range data {
		switch foundCounter {
		case 1:
			if r == 'h' || r == 'H' {
				foundCounter++
			}
		case 2:
			if r == 't' || r == 'T' {
				foundCounter++
			}
		case 3:
			if r == 'm' || r == 'M' {
				foundCounter++
			}
		case 4:
			if r == 'l' || r == 'L' {
				foundCounter++
			}
		case 5:
			if r == '>' || r == ' ' {
				// found "<html " or "<html>"
				return pos, nil
			}
		default:
			if r == '<' {
				pos = i
				foundCounter = 1
			} else {
				foundCounter = 0
			}
		}
		if i > 200 {
			return 0, errNoHTML // no <html> tag for the first 200 bytes
		}
		i++
	}
	return 0, errNoHTML
}

func HTMLIndexString(s string) (uint64, error) {
	var (
		foundCounter uint8
		pos          uint64
		i            uint64
	)
	for _, r := range s {
		switch foundCounter {
		case 1:
			if r == 'h' || r == 'H' {
				foundCounter++
			}
		case 2:
			if r == 't' || r == 'T' {
				foundCounter++
			}
		case 3:
			if r == 'm' || r == 'M' {
				foundCounter++
			}
		case 4:
			if r == 'l' || r == 'L' {
				foundCounter++
			}
		case 5:
			if r == '>' || r == ' ' {
				// found "<html " or "<html>"
				return pos, nil
			}
		default:
			if r == '<' {
				pos = i
				foundCounter = 1
			} else {
				foundCounter = 0
			}
		}
		if i > 200 {
			return 0, errNoHTML // no <html> tag for the first 200 runes
		}
		i++
	}
	return 0, errNoHTML
}

func HasHTMLTagString(s string) bool {
	var (
		foundCounter uint8
		i            uint64
	)
	for _, r := range s {
		switch foundCounter {
		case 1:
			if r == 'h' || r == 'H' {
				foundCounter++
			}
		case 2:
			if r == 't' || r == 'T' {
				foundCounter++
			}
		case 3:
			if r == 'm' || r == 'M' {
				foundCounter++
			}
		case 4:
			if r == 'l' || r == 'L' {
				foundCounter++
			}
		case 5:
			if r == '>' || r == ' ' {
				// found "<html " or "<html>"
				return true
			}
		default:
			if r == '<' {
				foundCounter = 1
			} else {
				foundCounter = 0
			}
		}
		if i > 200 {
			return false // no <html> tag for the first 200 runes
		}
		i++
	}
	return false
}

func HasScriptTagString(s string) bool {
	var (
		foundCounter uint8
		i            uint64
	)
	for _, r := range s {
		switch foundCounter {
		case 1:
			if r == 's' || r == 'S' {
				foundCounter++
			}
		case 2:
			if r == 'c' || r == 'C' {
				foundCounter++
			}
		case 3:
			if r == 'r' || r == 'R' {
				foundCounter++
			}
		case 4:
			if r == 'i' || r == 'I' {
				foundCounter++
			}
		case 5:
			if r == 'p' || r == 'P' {
				foundCounter++
			}
		case 6:
			if r == 't' || r == 'T' {
				foundCounter++
			}
		case 7:
			if r == '>' || r == ' ' {
				// found "<script " or "<script>
				return true
			}
		default:
			if r == '<' {
				foundCounter = 1
			} else {
				foundCounter = 0
			}
		}
		i++
	}
	return false
}

func GetHTMLTag(data []byte) ([]byte, error) {
	pos, err := HTMLIndex(data)
	if err != nil {
		return nil, err
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
	pos, err := HTMLIndexString(s)
	if err != nil {
		return "", err
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
