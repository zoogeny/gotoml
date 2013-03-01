package gotoml

import (
	"errors"
	"strconv"
	"time"
)

func (m TOMLMap) GetString(key string) (s string, e error) {
	exists := false
	s, exists = m[key]
	if !exists {
		e = errors.New("Attempt to get property that does not exist")
		return
	}
	return
}

func (m TOMLMap) GetBool(key string) (b bool, e error) {
	str, exists := m[key]
	if !exists {
		e = errors.New("Attempt to get property that does not exist")
		return
	}

	switch str {
	case "true":
		b = true
	case "false":
		b = false
	default:
		e = errors.New("Invalid value for bool")
	}
	return
}

func (m TOMLMap) GetInt64(key string) (i int64, e error) {
	str, exists := m[key]
	if !exists {
		e = errors.New("Attempt to get property that does not exist")
		return
	}

	i, e = strconv.ParseInt(str, 10, 64)
	return
}

func (m TOMLMap) GetFloat64(key string) (f float64, e error) {
	str, exists := m[key]
	if !exists {
		e = errors.New("Attempt to get property that does not exist")
		return
	}

	f, e = strconv.ParseFloat(str, 64)
	return
}

func (m TOMLMap) GetTime(key string) (t time.Time, e error) {
	str, exists := m[key]
	if !exists {
		e = errors.New("Attempt to get property that does not exist")
		return
	}

	t, e = time.Parse(time.RFC3339, str)
	return
}
