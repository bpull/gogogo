//Must be at the top of every running program
package main

//How we import libraries
import (
  "fmt"  // Used for Printing
  )

//Global var for how many lines you have
var how_Many int = 24;

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

//A function that will print out how many lines are left
func addSomeLines(line int){

        // All factor_prime(x) calls will print the prime factors of the line number
        // All fmt.Print statements will just provide the words for that line
        factor_prime(line)
        fmt.Print(" lines of text on the screen, ")
        factor_prime(line)
        fmt.Println(" lines of text.")
        fmt.Print("Print it out, stand up and shout, ")
        factor_prime(line+1)
        fmt.Println(" lines of text on the screen.\n")
}

//Our required main function
func main(){

    for i := 1; i < how_Many; i++{  //For loop for going through all of the lines

      addSomeLines(i) //Our function to print out our lines
    }
}
