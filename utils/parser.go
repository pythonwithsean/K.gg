package utils

import "strings"

func ParseCommand(line string) Command {
	switch cmd := strings.ToUpper(strings.Split(line, " ")[0]); cmd {
	case "GET":
		return GET
	case "PUT":
		return PUT
	case "DEL":
		return DEL
	case "LIST":
		return LIST
	default:
		return UNKNOWN
	}
}

func ParsePut(line string) (string, string, bool) {
	parts := strings.Split(line, " ")
	if len(parts) != 3 {
		return "", "", false
	}

	return parts[1], parts[2], true
}

func ParseGetOrDel(line string) (string, bool) {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		return "", false
	}

	return parts[1], true
}
