package main

import "testing"

func TestCompareStrings(t *testing.T) {
	t.Run("test1", simpleTest("машина", "шимана", true))
	t.Run("test2", simpleTest("кино", "оник", true))
	t.Run("test3", simpleTest("раззаразом", "зараза", false))
	t.Run("test4", simpleTest("кормик", "микрон", false))
}

func simpleTest(stringOne, stringTwo string, answer bool) func(t *testing.T) {
	return func(t *testing.T) {
		result := CompareStrings(stringOne, stringTwo)
		if result != answer {
			t.Errorf("something went wrong, expected: %v, got: %v", answer, result)
		}
	}
}
