package hob

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/mrb/hob"
)

func TestTimestamp(t *testing.T) {
	assert.T(t, len(hob.Timestamp()) > 23)
}
