package hob

import ()

type SetData map[string]bool

type Set struct {
	setData SetData
	Set     []string
}

func NewSet() (set *Set) {
	return &Set{
		setData: newSetData(),
		Set:     newSetSlice(),
	}
}

func newSetData() (setData SetData) {
	setData = make(SetData)
	return
}

func newSetSlice() (setSlice []string) {
	setSlice = make([]string, 0)
	return
}

func (set *Set) Add(value string) (ok bool) {
	set.setData[value] = true
	ok = true
	return
}

func (set *Set) Remove(value string) (ok bool) {
	delete(set.setData, value)
	ok = true
	return
}

func (set *Set) Test(value string) (ok bool) {
	ok = set.setData[value]
	return
}

func (set *Set) Clone() (clone *Set) {
	clone = &Set{
		setData: set.setData,
		Set:     set.Set,
	}
	return
}

func (set *Set) Union(oset *Set) (union *Set) {
	union = set.Clone()
	for value, _ := range oset.setData {
		union.setData[value] = true
	}
	return
}

func (set *Set) Intersection(oset *Set) (intersection *Set) {
	intersection = NewSet()

	if len(set.setData) > len(oset.setData) {
		for value, _ := range oset.setData {
			if ok := set.setData[value]; ok {
				intersection.setData[value] = true
			}
		}
	} else {
		for value, _ := range set.setData {
			if ok := oset.setData[value]; ok {
				intersection.setData[value] = true
			}
		}
	}

	return
}

func (set *Set) Slice() []string {
	set.Set = newSetSlice()

	for value, ok := range set.setData {
		if ok {
			set.Set = append(set.Set, value)
		}
	}

	return set.Set
}
