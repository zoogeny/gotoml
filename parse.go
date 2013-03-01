package gotoml

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func StripLineComment(in string) string {
	var (
		out              string          = in
		r                *strings.Reader = strings.NewReader(in)
		err              error           = nil
		escape           bool            = false
		quoteCount       int             = 0
		commentRuneIndex int             = -1
		totalRead        int             = 0
		ch               rune
		size             int
	)
	for err != io.EOF {
		escape = false
		ch, size, err = r.ReadRune()
		totalRead += size
		if ch == '\\' {
			escape = true
			ch, size, err = r.ReadRune()
			totalRead += size
		}
		switch {
		case ch == '"' && !escape:
			quoteCount += 1
		case ch == '#' && quoteCount%2 == 0:
			commentRuneIndex = totalRead - 1
			break
		}
	}
	if commentRuneIndex != -1 {
		out = in[:commentRuneIndex]
	}
	return out
}

func ParseKeyValue(keyValue string, tomlMap TOMLMap) error {
	index := strings.IndexRune(keyValue, '=')
	if index == -1 {
		return errors.New("Expect = in key value string")
	}
	key := strings.TrimSpace(keyValue[:index])
	value := strings.TrimSpace(keyValue[index+1:])

	switch {
	case value[0] == '"':
		end := len(value) - 1
		tomlMap[key] = value[1:end]
	default:
		tomlMap[key] = value
	}

	return nil
}

func OpenTOML(path string) (tomlMap TOMLMap, err error) {
	var (
		file   *os.File = nil
		prefix string   = ""
		line   string
	)
	tomlMap = make(TOMLMap)
	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()

	r := bufio.NewReader(file)
	line, err = Readln(r)
	for err == nil {
		stripped := StripLineComment(line)
		trimmed := strings.TrimSpace(stripped)
		switch {
		case trimmed[0] == '[':
			end := len(trimmed) - 1
			prefix = strings.TrimSpace(trimmed[1:end-1]) + "."
		case strings.ContainsAny(trimmed, "="):
			err = ParseKeyValue(prefix+trimmed, tomlMap)
		}

		line, err = Readln(r)
	}

	if err == io.EOF {
		err = nil
	}

	return
}
