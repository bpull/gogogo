//Must be at the top of every running program
package main

//How we import libraries
import (
  "fmt"        // Used for Printing
//  "os"         // Used for opening files
//  "bufio"      // Used to read in a file one string at a time
//  "io/ioutil"  // Used to read in names of files in a directory
//  "strings"    // Used to split a phrase up into seperate words
)

func channel_tester(c chan [2]int,putin int){
 holder := [2]int{putin,putin+1}
 c <- holder
}

func main(){
  c := make(chan [2]int)
  for i:=0;i < 4; i++{
    go channel_tester(c,i)
  }

  for i:= 0; i < 4; i++{
    holder:= <-c
    fmt.Println("i = ",i," and holders = ",holder[0],"and",holder[1])
  }

}
