//Must be at the top of every running program
package main

//How we import libraries
import (
  "fmt"       // Used for Printing
  "math/rand" // Used for the rand function
  "time"      // used for seeding rand with time
  )

//Global var for how many lines you have
var how_Many int = 100;

//Global var for how much we should iterate every time
var increment int = 0;

//A function that will print out how many lines are left
func addSomeLines(lines int){
  // If this is the first line say line, not liness
  if lines == 1{
    fmt.Println(lines,"line of text on the screen,",lines,"line of text.")
  } else { // Print out lines if not at 1
    fmt.Println(lines,"lines of text on the screen,",lines,"lines of text.")
  }
  fmt.Println("Print it out, stand up and shout,", lines+increment, "lines of text on the screen.\n")
}

//Our required main function
func main() {
  rand.Seed(time.Now().UTC().UnixNano()) //How we seed our rand with time

  increment = rand.Intn(10)+1 // Set how much we will increment from 1-10

  line_counter := 1 // What we will actually use to count the lines

  for i := 1; i < how_Many; i++{ //For loop for going through all of the lines

    addSomeLines(line_counter)   // Our function to print out our lines

    line_counter += increment    // Add our increment value to our line counter
  }
}
