package easy

import "bytes"

func replaceSpace(s string) string {
	buf := bytes.Buffer{}
	for _, i := range s {
		if string(i) == " " {
			buf.WriteString("%20")
		} else {
			buf.WriteRune(i)
		}
	}
	return buf.String()
}
