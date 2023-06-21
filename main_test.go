package main

import "testing"

func TestSum(t *testing.T) {
    result := sum(5, 10)
    expected := 15

    if result != expected {
        t.Errorf("Expected %d, but got %d", expected, result)
    }
}
