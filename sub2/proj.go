//Must be at the top of every running program
package main

//How we import libraries
import (
  "fmt"        // Used for Printing
  "os"         // Used for opening files
  "bufio"      // Used to read in a file one string at a time
  "io/ioutil"  // Used to read in names of files in a directory
  "strings"    // Used to split a phrase up into seperate words
)


/*
This function will be spanwed as a go routine(also known as a thread) every time it is called.
1) It will open up the dictionary that it is passed
2) It will loop through every word in the dictionary and see if it exist in the map(hash table) of the text file you submitted
3) For every word that matches between the dictionary and the map, the length of the word is added to a counting variable
4) The quality count that came from the matching words, and the name of the dictionary that was searched is passed back to main.
*/
func how_many_words(c chan [2]int,language string,phrase_map map[string]int,which_lang int){

  count := 0

  // 1
  f, _ := os.Open(language)
  scanner := bufio.NewScanner(f)
  scanner.Split(bufio.ScanWords)
  for scanner.Scan(){
    line := scanner.Text()

    // 2
    _, exist := phrase_map[line]

    // 3
    if exist == true {
        count += len(line)
    }
  }
  f.Close()

  // 4
  holder := [2]int{count, which_lang}
  c <- holder
}

/*
The main function
1) It validates that you entered enought arguments on the command line, and that you gave it a valid text file to read
2) Reads in the file to array, and then sets that array as a hash table for faster lookup times later
3) Reads in all of the dictionary file names in the dictionary directory
4) Initializes variables needed for the go routines
5) Loops through all of the dictionary names and spawns a thread that will count how many words from your text file is in that dictionary
6) It will collect all of the returned values from our go routines
7) It will loop through all of the returned values and print the language that had the highest quality word count returned
*/
func main(){

  // 1
  if len(os.Args) != 2 {
    panic("Not Enough Arguments!\n usage: go run phrasehashwhich_language.go <text file>")
  }
  phrase, error := ioutil.ReadFile(os.Args[1])
  if error != nil{
    panic("Invalid Argument!\n usage: go run phrasehashwhich_language.go <text file>")
  }

  // 2
  phrase_array := strings.Fields(string(phrase))
  phrase_map := make(map[string]int)
  for i:=0; i<len(phrase_array);i++{
      phrase_map[phrase_array[i]] = 1
  }

  // 3
  files,_ := ioutil.ReadDir("./dictionaries")

  // 4
  var langs_string [18]string
  var all_langs [18]int
  count := 0
  max := 0

  // A channel variable is a variable that acts as a pipe between a go routine(thread) and the main function
  c := make(chan [2]int)

  // 5
  for _,file := range files {
    langs_string[count]=file.Name()
    language := "dictionaries/" + file.Name()
    go how_many_words(c,language,phrase_map,count)
    count++
  }

  // 6
  for i:=0;i < len(files); i++{
    holder := <-c
    all_langs[holder[1]] = holder[0]
  }

  // 7
  for j := 0; j < len(files); j++{
    if all_langs[j] > max{
      max = all_langs[j]
    }
  }
  if max == 0{
    fmt.Println("We do not have a dictionary that matches this language")
  } else {
    for i:= 0; i < len(files); i++{
      if all_langs[i] == max{
        fmt.Println(langs_string[i], "is the language you are looking for")
      }
    }
  }
}
