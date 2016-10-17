//Must be at the top of every running program
package main

//How we import libraries
import (
  "fmt" // Used for Printing
  )

//Global var for how many beers you have
var how_Many int = 100;

//A function that will print out how many beers are left
func addSomeBeer(bottlesOfBeer int){
  // If this is the first beer say bottle, not bottles
  if bottlesOfBeer == 1{
    fmt.Println(bottlesOfBeer,"bottle of beer on the wall,",bottlesOfBeer,"bottle of beer.")
  } else { // Print out beers if not at 1
    fmt.Println(bottlesOfBeer,"bottles of beer on the wall,",bottlesOfBeer,"bottles of beer.")
  }
  fmt.Println("Take one down, pass it around,", bottlesOfBeer+1, "bottles of beer on the wall.\n")
}

//Our required main function
func main() {

  for i := 1; i < how_Many; i++{ //For loop for going through all of the beers
    addSomeBeer(i) //Our function to print out our beers
  }
}
