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
func TestJudgePassLineTrue(t *testing.T) {
	output := JudgePassLine(70)
	assert.Equal(t, true, output)
}
func TestJudgePassLineFail(t *testing.T) {
	output := JudgePassLine(50)
	assert.Equal(t, false, output)
}
