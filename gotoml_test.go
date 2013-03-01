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
	if val, err := tomlMap.GetString("stringkey"); val != "string value" || err != nil {
		c.Error("should parse string values")
	}

	ParseKeyValue("boolkey = true", tomlMap)
	if val, err := tomlMap.GetBool("boolkey"); val != true || err != nil {
		c.Error("should parse bool values")
	}

	ParseKeyValue("intkey = 123456", tomlMap)
	if val, err := tomlMap.GetInt64("intkey"); val != 123456 || err != nil {
		c.Error("should parse int values")
	}

	ParseKeyValue("floatkey = 12.34", tomlMap)
	if val, err := tomlMap.GetFloat64("floatkey"); val != 12.34 || err != nil {
		c.Error("should parse float values")
	}

	ParseKeyValue("datekey = 1979-05-27T07:32:00Z", tomlMap)
	expectedDate := time.Date(1979, 5, 27, 7, 32, 0, 0, time.UTC)
	if val, err := tomlMap.GetTime("datekey"); val != expectedDate || err != nil {
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

func TestExampleTOML(c *testing.T) {
	out, err := OpenTOML("test-data/example.toml")
	if err != nil {
		c.Error("error parsing test-data/example.toml")
	}

	if out["title"] != "TOML Example" {
		c.Error("expecting title as TOML Example")
	}

	if out["owner.name"] != "Tom Preston-Werner" {
		c.Error("expecting owner.name as Tom Preston-Werner")
	}

	expectedDate := time.Date(1979, 5, 27, 7, 32, 0, 0, time.UTC)
	if val, err := out.GetTime("owner.dob"); val != expectedDate || err != nil {
		c.Error("should parse ISO8601 dates")
	}

	if out["servers.alpha.ip"] != "10.0.0.1" {
		c.Error("should parse nested tags")
	}
}

func TestExampleHardTOML(c *testing.T) {
	out, err := OpenTOML("test-data/hard-example.toml")
	if err != nil {
		c.Error("error parsing test-data/hard-example.toml")
	}

	if out["the.test_string"] != "You'll hate me after this - #" {
		c.Error("should handle # within strings")
	}
}
