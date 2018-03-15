package main

import "testing"

func TestToken(t *testing.T) {
	if _, err := token([]string{"aa", "fefe"}); err == nil {
		t.Errorf("Expected error of type argument, but got %v", err)
	}

	if _, err := token([]string{"aa", "1"}); err != nil {
		t.Errorf("Expected token, but got %v", err)
	}

	if _, err := token([]string{"aa", "1111111"}); err == nil {
		t.Errorf("Expected error of big argument, but got %v", err)
	}
}
