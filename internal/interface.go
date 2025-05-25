package internal

type Parse interface {
	ParseProto(content string) []Pattern
}

type Parser struct{}

type Pattern struct {
	Name     string
	Type     string
	Line     int
	StartPos int
	EndPos   int
}

type stack []string

type block struct {
	kind string
}
