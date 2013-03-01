package gotoml

import (
	"strconv"
	"time"
)

func (m TOMLMap) GetString(key string) (s string, e error) {
	exists := false
	s, exists = m[key]
	if !exists {
		e = NewKeyNotFoundError(key, "string")
		return
	}
	return
}

func (m TOMLMap) GetBool(key string) (b bool, e error) {
	str, exists := m[key]
	if !exists {
		e = NewKeyNotFoundError(key, "bool")
		return
	}

	switch str {
	case "true":
		b = true
	case "false":
		b = false
	default:
		e = NewInvalidTypeError(key, str, "bool")
	}
	return
}

func (m TOMLMap) GetInt64(key string) (i int64, e error) {
	str, exists := m[key]
	if !exists {
		e = NewKeyNotFoundError(key, "int64")
		return
	}

	i, e = strconv.ParseInt(str, 10, 64)
	return
}

func (m TOMLMap) GetFloat64(key string) (f float64, e error) {
	str, exists := m[key]
	if !exists {
		e = NewKeyNotFoundError(key, "float64")
		return
	}

	f, e = strconv.ParseFloat(str, 64)
	return
}

func (m TOMLMap) GetTime(key string) (t time.Time, e error) {
	str, exists := m[key]
	if !exists {
		e = NewKeyNotFoundError(key, "time")
		return
	}

	t, e = time.Parse(time.RFC3339, str)
	return
}
