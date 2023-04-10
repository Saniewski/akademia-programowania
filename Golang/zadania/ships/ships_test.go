package ships

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExamplePoint_Add() {
	p := Point{X: 1, Y: 2}
	a := Point{X: 2, Y: 2}
	p = p.Add(a)
	fmt.Println(p)
	// Output: {3 4}
}

func TestShip_MoveTo(t *testing.T) {
	s := Ship{
		Point{X: 1, Y: 2},
		Point{X: 2, Y: 1},
	}
	p := Point{X: 5, Y: 5}
	s = s.MoveTo(p)
	assert.Equal(t, Point{X: 5, Y: 5}, s[0])
	assert.Equal(t, Point{X: 6, Y: 4}, s[1])
}
