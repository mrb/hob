package hob

import (
	"log"
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

	set.Add("Key3")
	set.Add("Key4")
	set.Add("Key5")
	set.Add("Key6")
	set.Add("Key7")

	log.Print(set.Slice())

	set2 := hob.NewSet()
	set2.Add("Key1")
	set2.Add("Okey1")
	set2.Add("Okey23")

	union := set.Union(set2)
	log.Print(union.Slice())

	intersection := set.Intersection(set2)
	log.Print(intersection.Slice())
}
