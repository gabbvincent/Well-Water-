//Name: Vincent
//Date: 4/16/2020
//Description: Well Water

package main

import "fmt"

func Wellwater(a, b float64)float64{

  var GallonsPerfoot float64 = a * b * .163 / 100
  var Water float64 = b * GallonsPerfoot

  return Water 
}

func main() {
  var a float64
  var b float64

  fmt.Println("Enter the radius of your well casing in inches")
  fmt.Scanln(&a)
  fmt.Println("Enter how deep your well is in feet.")
  fmt.Scanln(&b)
  fmt.Println("Your Well casing will hold",Wellwater(a, b),"Gallons of water.")
}