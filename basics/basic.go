// every go program is made up of packages
package main  // programs start running in package main.... 

// this program uses the packages with import paths fmt and math/rand
import (
  "fmt"
  "math/rand"   // have to provide the path directly to the import we want access to
)

func main() {
  // methods called that belong to packages outside this package begin with a capital letter
  // a name is exported if it begins with a capital letter, e.g. Pizza is exported but pizza would not be
  // when importing a package you can only reger to its exported names
  // any "unexported" names are not accessible from outside the pacakge

  // NOTE: rand needs a seed (rand.Seed) to provide random numbers      
  fmt.Println("My favourite number is", rand.Intn(10))
}
