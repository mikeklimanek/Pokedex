package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
    cases := []struct {
        input    string
        expected []string
    }{
        {
            input: "Hello World",
            expected: []string{
                "hello", 
                "world",
            },
        },
    }

    for _, cs := range cases {
        actual := cleanInput(cs.input)
        if len(actual) != len(cs.expected) {
            t.Errorf("Expected %v words, got %v", len(cs.expected), len(actual))
            continue
        }
        for i := range actual {
            actualWord := actual[i]
            expectedWord := cs.expected[i]
            if actualWord != expectedWord {
                t.Errorf("Expected %v, got %v", expectedWord, actualWord)
            }
        }
    }
}
