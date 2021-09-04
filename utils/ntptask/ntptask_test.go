package ntptask

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPrintTime(t *testing.T) {
	locTime := time.Now()
	testTime := printTime()
	assert.Equal(t, locTime.Local().Hour(), testTime.Local().Hour())
	assert.Equal(t, locTime.Local().Minute(), testTime.Local().Minute())
	assert.Equal(t, locTime.Local().Second(), testTime.Local().Second())
}
