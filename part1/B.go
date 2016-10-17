//Must be at the top of every running program
package main

//How we import libraries
import (
  "fmt" // Used for Printing
  )

//Global var for how many lines you have
var how_Many int = 100;

//A function that will print out how many lines are left
func addSomeLines(lines int){
  // If this is the first line say line, not liness
  if lines == 1{
    fmt.Println(lines,"line of text on the screen,",lines,"line of text.")
  } else { // Print out lines if not at 1
    fmt.Println(lines,"lines of text on the screen,",lines,"lines of text.")
  }
  fmt.Println("Print it out, stand up and shout,", lines+1, "lines of text on the screen.\n")

}

//Our required main function
func main() {
  for i := 1; i < how_Many; i++{ //For loop for going through all of the lines
    addSomeLines(i) //Our function to print out our lines
  }
}
