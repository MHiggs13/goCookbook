package main

import (
  "fmt"
  "image"
)

func main() {
  // Package image defines the Image interface
  // type Image interface { 
  //    ColorModel() color.Model
  //    Bounds() Rectangle
  //    At(x, y int) color.color
  //  }
  // NOTE: the rectangle return value of the Bounds method is actually an image.Rectangle,
  // as the declaration is inside the package image
  // the color.Color and color.Model types are also interfaces, but we'll ignore that,
  // by using the predefined implentations color.RGBA and color.RGBAModel
  // these interfaces are specifided by the image/color package
  m := image.NewRGBA(image.Rect(0, 0, 100, 100))
  fmt.Println(m.Bounds())
  fmt.Println(m.At(0, 0).RGBA())
}
