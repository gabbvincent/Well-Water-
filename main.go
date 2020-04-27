//Name: Vincent
//Date: 4/16/2020
//Description: Well Water

package main

import "fmt"

func Wellwater(a, b, c float64)float64{
  
//Declare variable for GallonsPerfoot as float 64
//Declare variable for Water as float 64 
  var GallonsPerfoot float64 = a * b * .163 / 100
  var Water float64 = b * GallonsPerfoot
  if Water / c > 62.5{

    fmt.Println("Your Well casing will hold",Water,"Gallons of water, pleanty to supply your household and no need to install another holding tank.")
    
  } else {

    fmt.Println("Your well holds", Water,"Gallons of water and that's not enough for your household, an additional tank is required to supply your household sufficiently.")
  }
  return Water 
}

func main() {

  //Declare variable for a as float64
  //Declare variable for b as float64
  var a float64
  var b float64
  var c float64
  
  //print message asking user to enter their radius as inches and scan for a
  fmt.Println("Enter the radius of your well casing in inches")
  fmt.Scanln(&a)
  
  //Print message asking userfor their well depth in feet and scan for b
  fmt.Println("Enter how deep your well is in feet.")
  fmt.Scanln(&b)

 //Ask user for the household size and scan for c
  fmt.Println("Enter in the amount of people in your household.")
  fmt.Scanln(&c)

 //call Wellwater(a, b, c)
  Wellwater(a, b, c)
}
//end
