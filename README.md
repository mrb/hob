## Hob: CRDT For Go

Go implementations of data structures from <a href="http://hal.inria.fr/docs/00/55/55/88/PDF/techreport.pdf">A comprehensive study of Convergent and Commutative Replicated Data Types</a>

### Examples

#### LWW-element-Set

```go
lwwset, err := hob.NewLWWSet("a")
lwwset.Add("Dude!")
lwwset.Remove("Dude!")
lwwset.Add("Other key")
json, _ := lwwset.JSON()
```

Produces JSON:

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

### Tests

`go test`

### Credits

hob is (c) Michael R. Bernstein, 2012

### License

hob is distributed under the MIT License, see `LICENSE` file for details.
