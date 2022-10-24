package main

import (
	"errors"
	"testing"
)

func TestUnpackString(t *testing.T) {
	t.Run("letters and digits", simpleTest(t, "a4bc2d5e", "aaaabccddddde"))
	t.Run("with only leters", simpleTest(t, "abcd", "abcd"))
	t.Run("Empty string", simpleTest(t, "", ""))
	t.Run("With escape sequence", simpleTest(t, `qwe\4\5`, `qwe45`))
	t.Run("With two digits after escape sequence", simpleTest(t, `qwe\45`, `qwe44444`))
	t.Run("With escape letter after escape sequence", simpleTest(t, `qwe\\5`, `qwe\\\\\`))
	t.Run("With unicode symbols", simpleTest(t, "а1б2в3", "аббввв"))
	t.Run("With 0 repeat", simpleTest(t, "a0b0c", "c"))
	t.Run("With multiple escape characters", simpleTest(t, `ab\\2`, `ab\\`))
	t.Run("With invalid escape sequence", simpleErrorTest(t, `\abcd`))
	t.Run("With invalid first letter", simpleErrorTest(t, `32ab`))
}

func simpleTest(t *testing.T, packedString, resultString string) func(t *testing.T) {
	return func(t *testing.T) {
		unpackedString, err := UnpackString(packedString)
		if err != nil {
			t.Errorf("something went wrong: %v", err)
			return
		}
		if resultString != unpackedString {
			t.Errorf("expected: %v, have: %v\n", resultString, unpackedString)
		}
	}
}

func simpleErrorTest(t *testing.T, packedString string) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := UnpackString(packedString)
		if !errors.Is(err, ErrIncorrectString) {
			t.Errorf("wrong type of err, expected: %v", ErrIncorrectString)
		}
	}
}
