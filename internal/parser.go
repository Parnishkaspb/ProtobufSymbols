package internal

import (
	"bufio"
	"fmt"
	"strings"
)

func (s *stack) push(val string) {
	*s = append(*s, val)
}

func (s *stack) pop() {
	if len(*s) > 0 {
		*s = (*s)[:len(*s)-1]
	}
}

func (s *stack) depth() int {
	return len(*s)
}

func (pat Pattern) String() string {
	return fmt.Sprintf("%s %s %d:%d-%d", pat.Name, pat.Type, pat.Line, pat.StartPos, pat.EndPos)
}

func (p *Parser) ParseProto(content string) []Pattern {
	keywords := map[string]string{
		"import":  "import",
		"service": "service",
		"rpc":     "rpc",
		"enum":    "enum",
		"message": "message",
	}

	var blockStack []block

	var patterns []Pattern
	scanner := bufio.NewScanner(strings.NewReader(content))
	lineNum := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineNum++
		trimmed := strings.TrimSpace(line)

		words := strings.Fields(trimmed)
		if len(words) == 0 {
			continue
		}

		if kind, ok := keywords[words[0]]; ok && (kind == "message" || kind == "enum" || kind == "service") {
			if len(blockStack) == 0 && len(words) > 1 {
				name := strings.Trim(words[1], `{(=;`)
				startIdx := strings.Index(line, name)
				if kind != "message" && kind != "enum" || len(blockStack) == 0 {
					patterns = append(patterns, Pattern{
						Name:     name,
						Type:     kind,
						Line:     lineNum,
						StartPos: startIdx + 1,
						EndPos:   startIdx + len(name) + 1,
					})
				}
				blockStack = append(blockStack, block{kind: kind})
				continue
			}
		}

		if words[0] == "import" && len(words) > 1 && len(blockStack) == 0 {
			importName := strings.Trim(words[1], `";`)
			startIdx := strings.Index(line, importName)
			patterns = append(patterns, Pattern{
				Name:     importName,
				Type:     "import",
				Line:     lineNum,
				StartPos: startIdx,
				EndPos:   startIdx + len(importName) + 2,
			})
			continue
		}

		if words[0] == "rpc" && len(words) > 1 && len(blockStack) == 1 && blockStack[0].kind == "service" {
			raw := words[1]
			name := raw
			if idx := strings.Index(raw, "("); idx != -1 {
				name = raw[:idx]
			}
			startIdx := strings.Index(line, name)
			patterns = append(patterns, Pattern{
				Name:     name,
				Type:     "method",
				Line:     lineNum,
				StartPos: startIdx + 1,
				EndPos:   startIdx + len(name),
			})
			continue
		}

		if strings.Contains(line, "}") && len(blockStack) > 0 {
			blockStack = blockStack[:len(blockStack)-1]
		}
	}

	return patterns
}
