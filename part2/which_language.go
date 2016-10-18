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


func how_many_words(c chan [2]int,language string,word string,which_lang int){

  count := 0
  phrase_array := strings.Fields(word)
  f, _ := os.Open(language)
  scanner := bufio.NewScanner(f)
  scanner.Split(bufio.ScanWords)
  for scanner.Scan(){
    line := scanner.Text()
    for i:=0; i < len(phrase_array);i++{
      if line == phrase_array[i]{
        count++
      }
    }
  }
  f.Close()

  holder := [2]int{count, which_lang}
  c <- holder
}

func main(){
  //Channel variable for keeping track of max word count in a dictionary and which dictionary is being used
  c := make(chan [2]int)
  //An array to hold all of the dictionary names


  var langs_string [18]string
  //An array to match the string value of the dictionary to an int
  var all_langs [18]int
  //A counter to keep track of how many dictionaries we have
  count := 0
  //Opens the dictionaries directory
  files,_ := ioutil.ReadDir("./dictionaries")
  //Loops through all the files in the dictionary directory
  for _,file := range files {

    //Adds the name of all of the files to the array
    langs_string[count]=file.Name()
    //Creates the string that will be sent to our function for opening the file
    language := "dictionaries/" + file.Name()
    //Calls our function to count the words in that file
    go how_many_words(c,language,"Hey hvorfor er det lige du tror du kender mig A shabi rasi hrakni Så jeg venter bare på du kommer med en kommentar Nogen burde sku bare klap i For der er mange som der misforstår, når vi går Render rundt og tror vi alle bistand får YO hold lige Kan vi ikke få nogle argumenter med noget hold i Så må i heller komme igen med noget bedre For jeg ser det mange steder Læser osse om det når jeg sidder, og tjekker det på tv men jeg gider ikke sige hvad de hedder Og de allerede ved at forberede nye ting Siger det kun fordi jeg kan være det bekendt Og de er allerede ved at forberede syge ting I må heller kom med et bedre argument Og jeg sidder og læser avisen, tænker hvorfor skriver de Nogle ting om og om igen For i denne verden som vi lever i Må vi heller' bare kom igen",count)
    //Increment the number of dictionaries we have started searching through
    count++
  }


  //Collects the values returned from our channels, which will be the max count of words in that dictionary
  for i:=0;i < 18; i++{
    holder := <-c
    all_langs[holder[1]] = holder[0]
  }

  max := 0
  //Loops through all of the maxes that we recieved and print's out the greatest one
  for j := 0; j < 18; j++{

    if all_langs[j] > max{
      max = all_langs[j]
    }
  }

  for i:= 0; i < 18; i++{
    if all_langs[i] == max{

      fmt.Println(langs_string[i], "is the language you are looking for")
    }
  }
}
