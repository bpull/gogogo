//Must be at the top of every running program
package main

//How we import libraries
import (
  "fmt"     // Used for Printing
  "os"      // Used for opening/writing to files
  "strconv" // Used for converting ints to strings
  )

//Global var for how many lines you have
var how_Many int = 5;

//A function that will print out how many lines are left to the screen
func addSomeLines(lines int){
  // If this is the first line say line, not liness
  if lines == 1{
    fmt.Println(lines,"line of text on the screen,",lines,"line of text.")
  } else { // Print out lines if not at 1
    fmt.Println(lines,"lines of text on the screen,",lines,"lines of text.")
  }
  fmt.Println("Print it out, stand up and shout,", lines+1, "lines of text on the screen.\n")

}

//A function that will print out how many beers are left to a file fo
func addSomeBeerToFile(fo *os.File,beer int){

  printBeer1 := strconv.Itoa(beer)   // Turns the number of beers you have from an int to a string
  printBeer2 := strconv.Itoa(beer+1) // Turns the number of beers you have + 1 from an int to a string

  // If this is the first beer say bottle, not bottles
  if beer == 1{
      fo.WriteString(printBeer1+" bottle of beer on the wall, "+printBeer1+" bottle of beer.\n")
  } else { // Print out bottles if not at 1
      fo.WriteString(printBeer1+" bottles of beer on the wall, "+printBeer1+" bottles of beer.\n")
  }
  fo.WriteString("Take one down, pass it around, "+printBeer2+" bottles of beer on the wall.\n\n")
}

func main(){

    fo,err := os.Create("out.txt") // Sets fo as file pointer and returns any errors to err

    if err != nil{ // If there is any errors then print them and stop the program
      panic(err)
    }

    for i := 1; i < how_Many; i++{ //For loop for going through all of the lines/beers

        addSomeLines(i)         //Our function to print out our lines
        addSomeBeerToFile(fo,i) //Our function to print out our beers
    }
}
