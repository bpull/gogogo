//Must be at the top of every running program
package main

//How we import libraries
import (
  "fmt"        // Used for Printing
  "os"         // Used for opening files
  "io/ioutil"  // Used to read in names of files in a directory
  "strings"    // Used to split a phrase up into seperate words
)

func check(e error) {
    if e != nil {
        panic("Invalid Arguments!\n usage: go run which_lang.go <text file>")
    }
}

func how_many_words(c chan [2]int,language string,phrase_array []string,which_lang int){

  lang, error := ioutil.ReadFile(language)
  check(error)

  lang_array := strings.Fields(string(lang))
  lang_map := make(map[string]int)

  for i:=0; i<len(lang_array);i++{
      lang_map[lang_array[i]] = 1
  }

  count := 0
  for i:=0; i<len(phrase_array);i++{
    _, exist := lang_map[phrase_array[i]]
    if exist == true {
        count++
    }
  }

  holder := [2]int{count, which_lang}
  c <- holder
}

func main(){
  //Channel variable for keeping track of max word count in a dictionary and which dictionary is being used
  c := make(chan [2]int)
  //An array to hold all of the dictionary names


  phrase, error := ioutil.ReadFile(os.Args[1])
  check(error)

  phrase_array := strings.Fields(string(phrase))

  var langs_string [18]string
  //An array to match the string value of the dictionary to an int
  var all_langs [18]int
  //A counter to keep track of how many dictionaries we have
  count := 0
  //Opens the dictionaries directory
  files,_ := ioutil.ReadDir("./dictionaries")

  //phrase := "Sólo tú le das brillo y amor a mi corazón. Sólo tú me miras con tanta pureza y belleza interior.Sólo tú llenas de alegría a mi alma.Sólo tú conoces mis secretos, deseos y anhelos.Sólo tú me transmites calor con tus manos.Sólo tú me das fuerzas para seguir viviendo.Sólo tú me hablas con tanta dulzura.Sólo tú me escuchas con tanta paciencia.Sólo tú me acaricias con tanta delicadeza.Sólo tú me besas con tanta ternura.Sólo tú me comprendes con tanta sutileza.Sólo tú haces que me sienta el ser más feliz del universo.Sólo tú me das tanta seguridad en mí misma.Sólo tú me calmas y tranquilizas cuando estoy mal.Sólo tú eres el que puede entrar en mi corazón y en mi ser.Sólo tú me abrigas cuando tengo frío.Sólo tú eres y serás el dueño de mi corazón.Sólo tú eres mi sueño hecho realidad.Sólo tú eres y serás el que me ama y me amará siempre.Sólo tú eres al que amo y amaré eternamente!"
  //Loops through all the files in the dictionary directory
  for _,file := range files {

    //Adds the name of all of the files to the array
    langs_string[count]=file.Name()
    //Creates the string that will be sent to our function for opening the file
    language := "dictionaries/" + file.Name()
    //Calls our function to count the words in that file
    go how_many_words(c,language,phrase_array,count)
    //Increment the number of dictionaries we have started searching through
    count++
  }

  //Collects the values returned from our channels, which will be the max count of words in that dictionary
  for i:=0;i < len(files); i++{
    holder := <-c
    all_langs[holder[1]] = holder[0]
  }

  max := 0
  //Loops through all of the maxes that we recieved and print's out the greatest one
  for j := 0; j < len(files); j++{

    if all_langs[j] > max{
      max = all_langs[j]
    }
  }

  for i:= 0; i < len(files); i++{
    if all_langs[i] == max{

      fmt.Println(langs_string[i], "is the language you are looking for")
    }
  }
}
