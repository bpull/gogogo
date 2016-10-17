//Must be at the top of every running program
package main

//How we import libraries
import (
  "fmt"       // Used for Printing
  "os"        // Used for opening/writing to files
  "strconv"   // Used for converting ints to strings
  "time"      // Used for sleep function
  "math/rand" // Used for the rand function

  )

//Global var for how many lines you have
var how_Many int = 5;
var increment int = 1;

//A function that will print off the prime factors of a number
func factor_prime(numb int){

    // A flag for determing if this is the first time printing a number for formatting
    var used int = 0

    // A flag for determing if you have hit the end without finding any prime factors
    var orignumb int = numb

    // Always print a ( regardless
    fmt.Print("(")

    // for loop that will start at 2 and check for primes
    for i := 2; i <= numb; i++{

        if numb%i == 0{ // if numb%i = 0 then i is a factor

            is_i_prime := true // flag for determing if i is prime

            // Run the for loop till i to determine if prime
            for j := 2; j < i; j++{

                if i%j == 0{ // If j evenly divides i then it is not prime and we skip it

                    is_i_prime = false // Set flag to false
                    break // Break out of for loop
                }
            }

            //If i is prime and i is not the original number then print it out
            if is_i_prime && i != orignumb {

                if used == 0{ // if this is the first time printing format correctly
                  fmt.Print(i)
                  used++
                }else{
                  fmt.Print("*",i)
                }

                numb = numb/i // divide numb by i so that we can start over on a smaller search space this time
                i = 1 // set i = 1 so that we can restart at 2 because i++ happens at the end of the loop
            }
        }
    }
    fmt.Print(")") // regardless of what was found print off closing )
}

//A function that will print off the prime factors of a number
func factor_prime_file(numb int, fo *os.File){

    // A flag for determing if this is the first time printing a number for formatting
    var used int = 0

    // A flag for determing if you have hit the end without finding any prime factors
    var orignumb int = numb

    // Always print a ( regardless
    fo.WriteString("(")

    // for loop that will start at 2 and check for primes
    for i := 2; i <= numb; i++{

        if numb%i == 0{ // if numb%i = 0 then i is a factor

            is_i_prime := true // flag for determing if i is prime

            // Run the for loop till i to determine if prime
            for j := 2; j < i; j++{

                if i%j == 0{ // If j evenly divides i then it is not prime and we skip it

                    is_i_prime = false // Set flag to false
                    break // Break out of for loop
                }
            }

            //If i is prime and i is not the original number then print it out
            if is_i_prime && i != orignumb {

                if used == 0{ // if this is the first time printing format correctly
                  //fmt.Print(i)
                  fo.WriteString(strconv.Itoa(i))
                  used++
                }else{
                  fo.WriteString("*"+strconv.Itoa(i))
                }

                numb = numb/i // divide numb by i so that we can start over on a smaller search space this time
                i = 1 // set i = 1 so that we can restart at 2 because i++ happens at the end of the loop
            }
        }
    }
    fo.WriteString(")") // regardless of what was found print off closing )
}

//A function that will print out how many lines are left to the screen
func addSomeLines(c chan int){ // c is a channel for passing ints to parent

    for line := 1; line < how_Many; line=line+increment{ //For loop for going through all of the lines

      // All factor_prime(x) calls will print the prime factors of the line number
      // All fmt.Print statements will just provide the words for that line
      factor_prime(line)
      fmt.Print(" lines of text on the screen, ")
      factor_prime(line)
      fmt.Println(" lines of text.")
      fmt.Print("Print it out, stand up and shout, ")
      factor_prime(line+increment)
      fmt.Println(" lines of text on the screen.\n")
      time.Sleep(time.Duration(1)*time.Second) //Sleeps the program for 1 second
  }

  fmt.Println("Screen finished") // Print to the screen that this function has finished

  c <- 1 // pass one back to parent process through channel c
}

//A function that will print out how many beers are left to a file fo
func addSomeBeerToFile(fo *os.File,c chan int){

  for beer := 1; beer < how_Many; beer=beer+increment{
    // All factor_prime(x) calls will print the prime factors of the line number
    // All fo.WriteString statements will just provide the words for that line
    factor_prime_file(beer,fo)
    fo.WriteString(" bottles of beer on the wall, ")
    factor_prime_file(beer,fo)
    fo.WriteString(" bottles of beer.\n")
    fo.WriteString("Take one down, pass it around, ")
    factor_prime_file(beer+increment,fo)
    fo.WriteString(" bottles of beer on the wall.\n\n")
  }

  fmt.Println("File finished") // Print to the screen that this function has finished

  c <- 1 // pass one back to parent process through channel c
}

func main(){
    fo,err := os.Create("out.txt") // Sets fo as file pointer and returns any errors to err

    if err != nil{ // If there is any errors then print them and stop the program
      panic(err)
    }

    rand.Seed(time.Now().UTC().UnixNano()) //How we seed our rand with time

    increment = rand.Intn(10)+1 // Set how much we will increment from 1-10

    // this makes a channel which is a special structure in go
    // that allows goroutines to pass information back to the parent
    // is very useful for figuring out when a go routine is done
    c := make(chan int)
    fmt.Print("How many lines? ") // Ask how many lines we want to print
    fmt.Scanln(&how_Many) // Set how_Many to the users input

    go addSomeLines(c) // starts a concurrent go routine that will run the addSomeLines function
    go addSomeBeerToFile(fo,c) // starts a concurrent go routine that will run the addSomeBeerToFile function

    x,y := <-c, <-c // how we wait for wach go routine to finish

    // Can't leave a hanging statement in go so check if they equal one, print an error if they don't
    if x != 1 || y != 1{
      fmt.Println("There was an error somewhere")
    }
}
