package utils

const ADDR = "localhost"

const PORT = "8080"

type Command int

const (
	UNKNOWN Command = iota
	GET
	PUT
	DEL
	LIST
)
