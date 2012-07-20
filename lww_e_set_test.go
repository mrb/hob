package hob

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/mrb/hob"
)

func TestNewLWWESet(t *testing.T) {
	_, err := hob.NewLWWESet("nah")
	assert.T(t, err == hob.ErrInvalidBias)
}

func setupLWWESet(t *testing.T) (lwwset *hob.LWWESet) {
	lwwset, err := hob.NewLWWESet("a")
	assert.T(t, lwwset != nil)
	assert.T(t, err == nil)
	return
}

func setupLWWESetWithData(t *testing.T) (lwwset *hob.LWWESet) {
	lwwset = setupLWWESet(t)

	err := lwwset.Add("Key1")
	assert.T(t, err == nil)

	err = lwwset.Remove("Key1")
	assert.T(t, err == nil)

	err = lwwset.Add("Key2")
	assert.T(t, err == nil)

	return
}

/*
- Bias: "a" / "r"
*/
func TestLWWESetBias(t *testing.T) {
	lwwset := setupLWWESetWithData(t)
	assert.T(t, lwwset.Bias == "a")

	rlwwset := setupLWWESetWithData(t)
	rlwwset.Bias = "r"
	assert.T(t, rlwwset.Bias == "r")

	// e added, not removed - true / true
	is_member, err := lwwset.Test("Key2")
	assert.T(t, err == nil)
	assert.T(t, is_member == true)

	is_member, err = rlwwset.Test("Key2")
	assert.T(t, err == nil)
	assert.T(t, is_member == true)

	// e added, removed, removed > added - false / false
	is_member, err = lwwset.Test("Key1")
	assert.T(t, err == nil)
	assert.T(t, is_member == false)

	is_member, err = rlwwset.Test("Key1")
	assert.T(t, err == nil)
	assert.T(t, is_member == false)

	// e added, removed, added -- added = removed - true / false
	err = lwwset.Add("Key1")
	assert.T(t, err == nil)
	lwwset.Data["Key1"].Add = lwwset.Data["Key1"].Remove
	assert.T(t, lwwset.Data["Key1"].Add == lwwset.Data["Key1"].Remove)

	is_member, err = lwwset.Test("Key1")
	assert.T(t, err == nil)
	assert.T(t, is_member == true)

	err = rlwwset.Add("Key1")
	assert.T(t, err == nil)
	rlwwset.Data["Key1"].Add = rlwwset.Data["Key1"].Remove
	assert.T(t, rlwwset.Data["Key1"].Add == rlwwset.Data["Key1"].Remove)

	is_member, err = rlwwset.Test("Key1")
	assert.T(t, err == nil)
	assert.T(t, is_member == false)

	// e added, removed, added again - true / true
	err = lwwset.Add("Key1")
	assert.T(t, err == nil)

	is_member, err = lwwset.Test("Key1")
	assert.T(t, err == nil)
	assert.T(t, is_member == true)

	err = rlwwset.Add("Key1")
	assert.T(t, err == nil)

	is_member, err = rlwwset.Test("Key1")
	assert.T(t, err == nil)
	assert.T(t, is_member == true)
}

func TestLWWJson(t *testing.T) {
	lwwset := setupLWWESetWithData(t)

	json, err := lwwset.JSON()

	assert.T(t, err == nil)
	assert.T(t, json != nil)

	//data, err := hob.ParseJson(json)
	//assert.T(t, err == nil)
	//assert.T(t, data != nil)
	//assert.T(t, data.(LWWESet).Type == "lww-e-set")
	//assert.T(t, data.(*hob.LWWESet).JSONData != nil)
	//assert.T(t, len(data.(*hob.LWWESet).JSONData) == 2)
}

func TestLWWMerge(t *testing.T) {
	lwwset := setupLWWESetWithData(t)
	olwwset := setupLWWESetWithData(t)

	err := olwwset.Remove("Key2")
	assert.T(t, err == nil)

	merged, err := lwwset.Merge(olwwset)
	assert.T(t, err == nil)

	assert.T(t, merged.Data["Key2"] != nil)
	assert.T(t, merged.Data["Key2"].Add != "")
	assert.T(t, merged.Data["Key2"].Remove != "")
	assert.T(t, merged.Data["Key2"].Remove == olwwset.Data["Key2"].Remove)
	assert.T(t, merged.Data["Key2"].Add == olwwset.Data["Key2"].Add)
}
