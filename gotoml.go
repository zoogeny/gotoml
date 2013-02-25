package gotoml

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

type TOMLMap map[string]interface{}

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
	case strings.HasPrefix(value, "\""):
		end := len(value) - 1
		tomlMap[key] = value[1:end]
	case value == "true":
		tomlMap[key] = true
	case value == "false":
		tomlMap[key] = false
	}

	//_,err := time.Parse(time.RFC3339, "1979-05-27T07:32:00Z")
	return nil
}

func PushDict(dictId string, tomlMap TOMLMap) (newMap TOMLMap, err error) {
	return newMap, err
}

func OpenTOML(path string) (tomlMap TOMLMap, err error) {
	var (
		file       *os.File = nil
		line       string
		currentMap TOMLMap
	)
	tomlMap = make(TOMLMap)
	currentMap = tomlMap
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
		case strings.HasPrefix(trimmed, "["):
			currentMap, err = PushDict(trimmed, tomlMap)
		case strings.ContainsAny(trimmed, "="):
			err = ParseKeyValue(trimmed, currentMap)
		}

		line, err = Readln(r)
	}

	if err == io.EOF {
		err = nil
	}

	return
}
