package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelloTom(t *testing.T) {
	output := returnTom()
	exceptOutput := "Tom"
	assert.Equal(t, output, exceptOutput)
}
func TestJudgePassLine(t *testing.T) {
	output := JudgePassLine(70)
	assert.Equal(t, true, output)
}
