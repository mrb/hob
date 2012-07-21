package hob

import (
	"log"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/mrb/hob"
)

func TestNewTwoPhaseSet(t *testing.T) {
	two_phase_set, err := hob.NewTwoPhaseSet()
	assert.T(t, err == nil)
	assert.T(t, two_phase_set != nil)

	two_phase_set.Add("Key1")
	two_phase_set.Add("Key2")
	two_phase_set.Remove("Key1")

	jsonb, err := two_phase_set.JSON()
	assert.T(t, err == nil)
	assert.T(t, jsonb != nil)

	jsons := string(jsonb)
	log.Print(jsons)
}
