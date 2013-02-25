package gotoml

import (
	"fmt"
	"testing"
	"time"
)

func TestStripLineComment(c *testing.T) {
	out := StripLineComment("should remain # should strip")
	expected := "should remain "
	if out != expected {
		fmt.Println(out)
		c.Error("should strip from # to end of string")
	}

	out = StripLineComment("key = \"\\\"contains\\\" # and \\\" should remain\" # should strip")
	expected = "key = \"\\\"contains\\\" # and \\\" should remain\" "
	if out != expected {
		fmt.Println(out)
		c.Error("should not strip # from within strings")
	}
}

func TestParseKeyValue(c *testing.T) {
	tomlMap := TOMLMap{"foo": "bar"}

	ParseKeyValue("stringkey = \"string value\"", tomlMap)
	if tomlMap["stringkey"] != "string value" {
		c.Error("should parse string values")
	}

	ParseKeyValue("intkey = 123456", tomlMap)
	if tomlMap["intkey"] != 123456 {
		c.Error("should parse int values")
	}

	ParseKeyValue("floatkey = 12.34", tomlMap)
	if tomlMap["floatkey"] != 12.34 {
		c.Error("should parse float values")
	}

	ParseKeyValue("boolkey = true", tomlMap)
	if tomlMap["boolkey"] != true {
		c.Error("should parse bool values")
	}

	ParseKeyValue("datekey = 1979-05-27T07:32:00Z", tomlMap)
	expectedDate := time.Date(1979, 5, 27, 7, 32, 0, 0, time.UTC)
	if tomlMap["datekey"] != expectedDate {
		c.Error("should parse ISO8601 dates")
	}
}

func TestSimpleTOML(c *testing.T) {
	out, err := OpenTOML("test-data/simple.toml")
	if err != nil {
		c.Error("error parsing test-data/simple.toml")
	}

	if out["title"] != "Simple TOML" {
		c.Error("expected title as Simple TOML")
	}
}
