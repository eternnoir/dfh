package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParseStartTime(t *testing.T) {
	assert := assert.New(t)
	result, err := ParseStartTime("240h", "", time.Now())
	if err != nil {
		assert.Fail(err.Error())
		return
	}
	assert.Equal(time.Now().AddDate(0, 0, -10).Day(), result.Day())
}
