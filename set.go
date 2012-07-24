/*
This is extremely naive but it's my first shot at playing with extending
types in Go. What I actually want is a data structure that wraps the
map[string]bool type, and adds Add, Remove, and Test methods, in addition
to Union, Intersection, etc. There should be an easy, cached interface to
an array of strings which would represent the actual "set" of data, since
the map of string:bool is a hack.
*/
package hob

import ()

type Set map[string]bool

func NewSet() (set Set) {
	set = make(Set)
	return
}

func (set Set) Add(value string) (ok bool) {
	set[value] = true
	ok = true
	return
}

func (set Set) Remove(value string) (ok bool) {
	set[value] = false
	ok = true
	return
}

func (set Set) Test(value string) (ok bool) {
	ok = set[value]
	return
}
