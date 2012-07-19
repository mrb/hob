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
* e added, not removed - true / true
* e added, removed, removed > added - false / false
* e added, removed, removed = added - true / false
* e added, removed, added again - true / true
*/
func TestLWWESetBias(t *testing.T) {
	lwwset := setupLWWESetWithData(t)

	is_member, err := lwwset.Test("Key2")
	assert.T(t, err == nil)
	assert.T(t, is_member == true)

	is_member, err = lwwset.Test("NOOOOOOO")
	assert.T(t, err == nil)
	assert.T(t, is_member == false)

	err = lwwset.Remove("Key1")
	assert.T(t, err == nil)
	assert.T(t, is_member == false)
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
