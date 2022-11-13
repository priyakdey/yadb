package main

import (
    "testing"
)

func TestGreet(t *testing.T) {
    expected := "Hello"
    actual := greet()

    if actual != expected {
        t.Fail()
    }
}
