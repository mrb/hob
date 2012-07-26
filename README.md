[![Build Status](https://secure.travis-ci.org/mrb/hob.png?branch=master)](http://travis-ci.org/mrb/hob)

## Hob: CRDT For Go

Go implementations of data structures from <a href="http://hal.inria.fr/docs/00/55/55/88/PDF/techreport.pdf">A comprehensive study of Convergent and Commutative Replicated Data Types</a>

#### This is pre-release, experimental software

### Examples

#### Two-Phase-Set

```go
two_phase_set, _ := hob.NewTwoPhaseSet()
two_phase_set.Add("I'm in the add set")
two_phase_set.Add("I'm also in the add set")
two_phase_set.Remove("I'm in the add set") // and in the remove set
json, _ := two_phase_set.JSON()
```

Produces:

```json
{
  "type":"2p-set",
  "a": ["I'm in the add set","I'm also in the add set"],
  "r": ["I'm in the add set"]
}
```

#### LWW-element-Set

```go
lwwset, _ := hob.NewLWWSet("a")
lwwset.Add("Dude!")
lwwset.Remove("Dude!")
lwwset.Add("Other key")
json, _ := lwwset.JSON()
```

Produces:

```json
{
  "type":"lww-e-set",
  "bias":"a",
  "e":[
    ["Dude!","2012-07-16T00:42:05.146259Z","2012-07-16T00:42:05.146263Z"],
    ["Other key","2012-07-16T00:42:05.146267Z",""]
  ]
}
```


### Prior Art

A few Open Source implementations of these data structures exist. Hob conforms to the same JSON format as:

* Reid Draper *Knockbox* (<a href="https://github.com/reiddraper/knockbox">Clojure Implementation - Github Repo</a>)
* Kyle Kingsbury *Meangirls* (<a href="https://github.com/aphyr/meangirls">Ruby Implementation - Github Repo</a>)

### TODO

A lot! This library lacks a lot at the moment, including JSON decoding. I'm trying to figure out the best idiomatic Go way to handle parsing multiple data types with the same code.

### Tests

`go test`

### Credits

hob is (c) Michael R. Bernstein, 2012

### License

hob is distributed under the MIT License, see `LICENSE` file for details.
