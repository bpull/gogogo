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
func addSomeLines(c chan int){ // c is a channel for passing ints to parent

    for line := 1; line < how_Many; line++{ //For loop for going through all of the lines

      // If this is the first line say line, not liness
      if line == 1{
        fmt.Println(line,"line of text on the screen,",line,"line of text.")
      } else { // Print out lines if not at 1
        fmt.Println(line,"lines of text on the screen,",line,"lines of text.")
      }
      fmt.Println("Print it out, stand up and shout,",line+1,"lines of text on the screen.\n")
  }

  fmt.Println("Screen finished") // Print to the screen that this function has finished

  c <- 1 // pass one back to parent process through channel c
}

//A function that will print out how many beers are left to a file fo
func addSomeBeerToFile(fo *os.File,c chan int){

  for beer := 1; beer < how_Many; beer++{
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

  fmt.Println("File finished") // Print to the screen that this function has finished

  c <- 1 // pass one back to parent process through channel c
}

func main(){
    fo,err := os.Create("out.txt") // Sets fo as file pointer and returns any errors to err

    if err != nil{ // If there is any errors then print them and stop the program
      panic(err)
    }

    // this makes a channel which is a special structure in go
    // that allows goroutines to pass information back to the parent
    // is very useful for figuring out when a go routine is done
    c := make(chan int)

    go addSomeLines(c) // starts a concurrent go routine that will run the addSomeLines function
    go addSomeBeerToFile(fo,c) // starts a concurrent go routine that will run the addSomeBeerToFile function

    x,y := <-c, <-c // how we wait for wach go routine to finish

    // Can't leave a hanging statement in go so check if they equal one, print an error if they don't
    if x != 1 || y != 1{
      fmt.Println("There was an error somewhere")
    }
}
