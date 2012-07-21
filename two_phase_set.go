package hob

import (
	"encoding/json"
)

type TwoPhaseSet struct {
	Type  string            `json:"type"`
	A     map[string]string `json:"-"`
	R     map[string]string `json:"-"`
	JSONA []string          `json:"a"`
	JSONR []string          `json:"r"`
}

func NewTwoPhaseSet() (twoPhaseSet *TwoPhaseSet, err error) {
	add := make(map[string]string)
	remove := make(map[string]string)

	twoPhaseSet = &TwoPhaseSet{
		Type: "2p-set",
		A:    add,
		R:    remove,
	}
	return
}

func (twoPhaseSet *TwoPhaseSet) Add(value string) (err error) {
	twoPhaseSet.A[value] = ""
	return
}

func (twoPhaseSet *TwoPhaseSet) Remove(value string) (err error) {
	twoPhaseSet.R[value] = ""
	return
}

func (twoPhaseSet *TwoPhaseSet) Test(value string) (is_member bool, err error) {
	if _, ok := twoPhaseSet.R[value]; ok {
		is_member = false
		return
	}

	if _, ok := twoPhaseSet.A[value]; ok {
		is_member = true
		return
	}
	return
}

func (twoPhaseSet *TwoPhaseSet) JSON() (json_bytes []byte, err error) {
	for k, _ := range twoPhaseSet.A {
		twoPhaseSet.JSONA = append(twoPhaseSet.JSONA, k)
	}

	for k, _ := range twoPhaseSet.R {
		twoPhaseSet.JSONR = append(twoPhaseSet.JSONR, k)
	}

	json_bytes, err = json.Marshal(twoPhaseSet)
	return
}

func (twoPhaseSet *TwoPhaseSet) Clone() (clone *TwoPhaseSet, err error) {
	clone = &TwoPhaseSet{
		Type: twoPhaseSet.Type,
		A:    twoPhaseSet.A,
		R:    twoPhaseSet.R,
	}
	return
}

func (twoPhaseSet *TwoPhaseSet) Merge(oTwoPhaseSet *TwoPhaseSet) (merged_set *TwoPhaseSet, err error) {
	return
}
