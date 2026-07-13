package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare1(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(81, square(9), "square(9) should be 81")
	/*rst := square(9)
	if rst != 81 {
		t.Errorf("square(9) should be 81 but returns %d", rst)
	}*/
}
