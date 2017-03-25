package main

import (
	"fmt"
	"testing"
)

func TestReadChallenges(t *testing.T) {
	var cBlob = []byte(`[
	{"slug": "c1", "version": 4, "author": "cmdchallenge", "description": "foo", "example": "bar", "expected_output": {"lines": ["herp", "derp"]}},
	{"slug": "c2", "version": 4, "author": "cmdchallenge", "description": "foo", "example": "bar", "expected_output": {"lines": ["hello world"]}},
	{"slug": "c2", "author": "cmdchallenge", "description": "foo", "example": "bar", "expected_output": {"order": false, "lines": ["hello world"]}},
	{"slug": "c2", "version": 4, "author": "cmdchallenge", "description": "foo", "example": "bar"}
	]`)
	challenges, err := readChallenges(cBlob)
	if err != nil {
		t.Error("Expected no error, got ", err)
	}
	fmt.Printf("%#v", challenges)
}
