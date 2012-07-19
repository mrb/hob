package hob

import (
	"encoding/json"
	"errors"
	"time"
)

var (
	ErrInvalidBias = errors.New("invalid bias, must be a or r")
)

type Pair struct {
	Add    string `json:"add"`
	Remove string `json:"remove"`
}

type LWWESet struct {
	Type     string           `json:"type"`
	Bias     string           `json:"bias"`
	Data     map[string]*Pair `json:"-"`
	JSONData [][3]string      `json:"e"`
}

func NewLWWESet(bias string) (lwwset *LWWESet, err error) {
	if bias != "a" && bias != "r" {
		err = ErrInvalidBias
	}

	data := make(map[string]*Pair)

	lwwset = &LWWESet{
		Type: "lww-e-set",
		Bias: bias,
		Data: data,
	}
	return
}

func (lwwset *LWWESet) Add(value string) (err error) {
	data := lwwset.Data

	if pair, ok := lwwset.Data[value]; ok {
		pair.Add = Timestamp()
		return
	}

	data[value] = &Pair{
		Add: Timestamp(),
	}
	return
}

func (lwwset *LWWESet) Remove(value string) (err error) {
	data := lwwset.Data

	if pair, ok := lwwset.Data[value]; ok {
		pair.Remove = Timestamp()
		return
	}

	data[value] = &Pair{
		Remove: Timestamp(),
	}
	return
}

func (lwwset *LWWESet) Test(value string) (is_member bool, err error) {
	if pair, ok := lwwset.Data[value]; ok {
		if remove := pair.Remove; remove != "" {
			remove_time, err := time.Parse(time.RFC3339, pair.Remove)
			if err != nil {
				return false, err
			}

			add_time, err := time.Parse(time.RFC3339, pair.Add)
			if err != nil {
				return false, err
			}

			bias := lwwset.Bias

			if bias == "a" {
				if add_time.Before(remove_time) == false {
					is_member = true
				} else {
					is_member = false
				}
			} else {
				if add_time.After(remove_time) == true {
					is_member = true
				} else {
					is_member = false
				}
			}

			return is_member, err
		}

		is_member = true
	} else {
		is_member = false
	}
	return is_member, nil
}

func (lwwset *LWWESet) ToSet() (keys []string, err error) {
	return
}

func (lwwset *LWWESet) Clone() (clone *LWWESet, err error) {
	clone = &LWWESet{
		Type: lwwset.Type,
		Bias: lwwset.Bias,
		Data: lwwset.Data,
	}
	return
}

func (lwwset *LWWESet) Merge(olwwset *LWWESet) (merged_set *LWWESet, err error) {
	merged_set, err = lwwset.Clone()
	if err != nil {
		return nil, err
	}

	for k, v := range olwwset.Data {
		if pair, ok := lwwset.Data[k]; ok {
			merged, err := pair.merge(v)
			if err != nil {
				return nil, err
			}
			merged_set.Data[k] = merged
		} else {
			merged_set.Data[k] = v
		}
	}

	return
}

func (lwwset *LWWESet) JSON() (json_bytes []byte, err error) {
	for k, v := range lwwset.Data {
		inner := [3]string{k, v.Add, v.Remove}
		lwwset.JSONData = append(lwwset.JSONData, inner)
	}

	json_bytes, err = json.Marshal(lwwset)
	return
}

func newPair(value string) (p *Pair) {
	p = &Pair{
		Add:    "",
		Remove: "",
	}
	return
}

func (p *Pair) merge(op *Pair) (merged *Pair, err error) {
	merged = &Pair{
		Add:    "",
		Remove: "",
	}

	add, err := time.Parse(time.RFC3339, p.Add)
	if err != nil {
		return nil, err
	}

	oadd, err := time.Parse(time.RFC3339, op.Add)
	if err != nil {
		return nil, err
	}

	if add.After(oadd) == true {
		merged.Add = p.Add
	} else {
		merged.Add = op.Add
	}

	var remove, oremove time.Time

	if p.Remove != "" {
		remove, _ = time.Parse(time.RFC3339, p.Remove)
	} else {
		if op.Remove == "" {
			return
		} else {
			merged.Remove = op.Remove
			return
		}
	}

	if op.Remove != "" {
		oremove, _ = time.Parse(time.RFC3339, op.Remove)
	} else {
		if p.Remove == "" {
			return
		} else {
			merged.Remove = p.Remove
			return
		}
	}

	if remove.After(oremove) != true {
		merged.Remove = p.Remove
	} else {
		merged.Remove = op.Remove
	}

	return
}
