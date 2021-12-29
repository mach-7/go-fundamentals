package main

import (
	"math"
	"testing"
)

func TestGetName(t *testing.T) {
	s := Square{
		sideLength: 6.6,
		name:       "MySquare",
	}

	name := getName(s)
	if name != "MySquare" {
		t.Errorf("Expected MySquare, but received %s", name)
	}
}

func TestGetArea(t *testing.T) {
	c := Circle{
		radius: 5.0,
		name:   "MyCircle",
	}

	area := getArea(c)
	expected := c.radius * c.radius * math.Pi
	if area != expected {
		t.Errorf("Expected Area %f, but got %f", expected, area)
	}
}
