package hob

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/mrb/hob"
)

func TestTimestamp(t *testing.T) {
	assert.T(t, len(hob.Timestamp()) > 23)
}

func TestNewLWWSet(t *testing.T) {
	_, err := hob.NewLWWSet("nah")
	assert.T(t, err == hob.ErrInvalidBias)
}

func setupLWWSet(t *testing.T) (lwwset *hob.LWWSet) {
	lwwset, err := hob.NewLWWSet("a")
	assert.T(t, lwwset != nil)
	assert.T(t, err == nil)
	return
}

func setupLWWSetWithData(t *testing.T) (lwwset *hob.LWWSet) {
	lwwset = setupLWWSet(t)

	err := lwwset.Add("Key1")
	assert.T(t, err == nil)

	err = lwwset.Remove("Key1")
	assert.T(t, err == nil)
	err = lwwset.Add("Key2")
	assert.T(t, err == nil)

	return
}

func TestLWWSetTest(t *testing.T) {
	lwwset := setupLWWSetWithData(t)

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
	lwwset := setupLWWSetWithData(t)

	json, err := lwwset.JSON()

	assert.T(t, err == nil)
	assert.T(t, json != nil)

	//data, err := hob.ParseJson(json)
	//assert.T(t, err == nil)
	//assert.T(t, data != nil)
	//assert.T(t, data.(LWWSet).Type == "lww-e-set")
	//assert.T(t, data.(*hob.LWWSet).JSONData != nil)
	//assert.T(t, len(data.(*hob.LWWSet).JSONData) == 2)
}

func TestLWWMerge(t *testing.T) {
	lwwset := setupLWWSetWithData(t)
	olwwset := setupLWWSetWithData(t)

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
