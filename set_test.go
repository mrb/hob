package hob

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/mrb/hob"
)

func TestSet(t *testing.T) {
	set := hob.NewSet()

	set.Add("Key1")

	key1 := set.Test("Key1")
	key2 := set.Test("Key2")

	assert.T(t, key1 == true)
	assert.T(t, key2 == false)

	set.Add("Key2")
	key2 = set.Test("Key2")
	assert.T(t, key2 == true)

	set.Remove("Key2")
	key2 = set.Test("Key2")
	assert.T(t, key2 == false)
}
