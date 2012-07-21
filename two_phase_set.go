package hob

import (
	"encoding/json"

//"errors"
)

type TwoPhaseSet struct {
	Type  string            `json:"type"`
	A     map[string]string `json:"-"`
	R     map[string]string `json:"-"`
	JSONA []string          `json:"a"`
	JSONR []string          `json:"r"`
}

func NewTwoPhaseSet() (two_phase_set *TwoPhaseSet, err error) {
	add := make(map[string]string)
	remove := make(map[string]string)

	two_phase_set = &TwoPhaseSet{
		Type: "2p-set",
		A:    add,
		R:    remove,
	}
	return
}

func (two_phase_set *TwoPhaseSet) Add(value string) (err error) {
	two_phase_set.A[value] = ""
	return
}

func (two_phase_set *TwoPhaseSet) Remove(value string) (err error) {
	two_phase_set.R[value] = ""
	return
}

func (two_phase_set *TwoPhaseSet) Test(value string) (is_member bool, err error) {
	return
}

func (two_phase_set *TwoPhaseSet) JSON() (json_bytes []byte, err error) {
	for k, _ := range two_phase_set.A {
		two_phase_set.JSONA = append(two_phase_set.JSONA, k)
	}

	for k, _ := range two_phase_set.R {
		two_phase_set.JSONR = append(two_phase_set.JSONR, k)
	}

	json_bytes, err = json.Marshal(two_phase_set)
	return
}

func (two_phase_set *TwoPhaseSet) Clone() (clone *TwoPhaseSet, err error) {
	clone = &TwoPhaseSet{
		Type: two_phase_set.Type,
		A:    two_phase_set.A,
		R:    two_phase_set.R,
	}
	return
}

func (two_phase_set *TwoPhaseSet) Merge(otwo_phase_set *TwoPhaseSet) (merged_set *TwoPhaseSet, err error) {
	return
}
