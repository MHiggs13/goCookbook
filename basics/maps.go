package main

import (
  "fmt"
  )

type Vertex struct {
  Lat, Long float64
}

var m map[string]Vertex
var m4 map[string]int

func mapLiterals() {
  // map literals are like struct literals but the keys are required
  var m2 = map[string]Vertex{
    "Bell Labs": Vertex{
      40.68433, -74.39967,
    },
    "Google": Vertex{
      37.42202, -122.08408,
    },
  }
  // if the top-level type is just a name, you can omit it form the elements of the literal
  var m3 = map[string]Vertex{
    "Bell Labs": { 40.68433, -74.39967},
    "Google": {37.42202, -122.08408},
  }

  fmt.Println(m2)
  fmt.Println(m3)
}

func insertingMaps() {
  // m[key] = elem    update
  m4["Answer"] = 42

}

func retrieveElement() {
  // elem = m[key]      get
  fmt.Println(m["Answer"])
}

func deleteElement() {
  // delete(m, key)
  delete(m, Answer)
}

func checkKeyPresent() {
  // test that key is present with a two-value assignment
  // elem, ok = m[key]
  // if key is in m, ok is true, otherwise ok is false
  // if key is not in the map, then elem is the zero value for the map's element type
  // if elem or ok have not been delcared yet you could use a short declaration form
  v, ok := m["Answer"]
}

func main() {
  // a map maps keys to values
  // the zero value of  amap is nil. A nil map has no keys, nor can keys be added
  // the make function returns a map of the given type, initalized and ready for use

  // make a map with keys of type string and with values of type Vertex
  m = make(map[string]Vertex)
  m["Bell Labs"] = Vertex{
    40.68433, -74.39967,
  }
  fmt.Println(m["Bell Labs"])

  m4 = make(map[string]int)
  //m := make(map[string]int)

}
