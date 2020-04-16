//Name: Vincent
//Date: 4/16/2020
//Description: Well Water

package main

import "fmt"

func Wellwater(a, b float64)float64{
  
//Declare variable for GallonsPerfoot as float 64
//Declare variable for Water as float 64 
  var GallonsPerfoot float64 = a * b * .163 / 100
  var Water float64 = b * GallonsPerfoot

  return Water 
}

func main() {
  //Declare variable for a as float64
  //Declare variable for b as float64
  var a float64
  var b float64
  
  //print message asking user to enter tjeir radius as inches and scan for a
  fmt.Println("Enter the radius of your well casing in inches")
  fmt.Scanln(&a)
  
  //Print message asking userfor their well depth in feet and scan for b
  fmt.Println("Enter how deep your well is in feet.")
  fmt.Scanln(&b)
  
  //Print message telling user the amount of water their well will hold by call Wellwater(a, b)
  fmt.Println("Your Well casing will hold",Wellwater(a, b),"Gallons of water.")
}
//end
