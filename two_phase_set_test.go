package hob

import (
	"log"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/mrb/hob"
)

func setupTwoPhaseSetWithData(t *testing.T) (twoPhaseSet *hob.TwoPhaseSet) {
	twoPhaseSet, err := hob.NewTwoPhaseSet()
	assert.T(t, err == nil)
	assert.T(t, twoPhaseSet != nil)

	twoPhaseSet.Add("Key1")
	twoPhaseSet.Add("Key2")
	twoPhaseSet.Remove("Key1")
	return
}

func TestNewTwoPhaseSet(t *testing.T) {
	twoPhaseSet := setupTwoPhaseSetWithData(t)
	assert.T(t, twoPhaseSet != nil)

	jsonb, err := twoPhaseSet.JSON()
	assert.T(t, err == nil)
	assert.T(t, jsonb != nil)

	jsons := string(jsonb)
	log.Print(jsons)
}

func TestTwoPhaseSetTest(t *testing.T) {
	twoPhaseSet := setupTwoPhaseSetWithData(t)
	assert.T(t, twoPhaseSet != nil)

	is_member, err := twoPhaseSet.Test("Key1")
	assert.T(t, err == nil)
	assert.T(t, is_member == false)

	is_member, err = twoPhaseSet.Test("Key2")
	assert.T(t, err == nil)
	assert.T(t, is_member == true)

}
