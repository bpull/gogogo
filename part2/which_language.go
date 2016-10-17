//Must be at the top of every running program
package main

//How we import libraries
import (
  "fmt"       // Used for Printing
  "os"        // Used for opening/writing to files
  "bufio"
  "io/ioutil"
  //"strconv"   // Used for converting ints to strings
  //"time"      // Used for sleep function
  //"math/rand" // Used for the rand function
  )

func how_many_words(c chan int,which chan string,language string,word string){

    f, _ := os.Open(language)
    count := 0
    scanner := bufio.NewScanner(f)
    scanner.Split(bufio.ScanWords)
    for scanner.Scan(){
        line := scanner.Text()
        if line == word{
            count++
        }
    }
    if count == 1{
        fmt.Println("The word",word,"showed up",count,"time in the",language,"dictionary")
    }
    if count > 1{
        fmt.Println("The word",word,"showed up",count,"times in the",language,"dictionary")
    }
    c <- count
    which <- language
}

func main(){
    //amer_english := "american-english"
    //bokmaal := "bokmaal"
    //brazilian := "brazilian"
    //brit_english := "brit_english"
    //bulgarian := "bulgarian"
    //catala := "catala"
    c := make(chan int)
    word := make(chan string)
    var all_langs [18]int
    var langs_string [18]string
    count := 0
    files,_ := ioutil.ReadDir("./dictionaries")
    for _,file := range files {

        langs_string[count]=file.Name()
        language := "dictionaries/" + file.Name()
        //fmt.Println(language)
        go how_many_words(c,word,language,"zygote")

        count++
    }
    fmt.Println(count)

    //go how_many_words(c,amer_english,"zygote")
    for i:=0;i < 18; i++{
        all_langs[i] = <-c
    }
    //a,b,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s := <-c,<-c,<-c,<-c,<-c,<-c,<-c,<-c,<-c,<-c,<-c,<-c,<-c,<-c,<-c,<-c,<-c,<-c

    for j := 0; j < 18; j++{
        if all_langs[j] > 0{
            fmt.Println(langs_string[j],"is the winner")
        }
    }

    //fmt.Println(x)
    /*
    if a != 1{
        fmt.Println("error in return")
    }
    if b != 1{
        fmt.Println("error in return")
    }
    if d != 1{
        fmt.Println("error in return")
    }
    if e != 1{
        fmt.Println("error in return")
    }
    if f != 1{
        fmt.Println("error in return")
    }
    if g != 1{
        fmt.Println("error in return")
    }
    if h != 1{
        fmt.Println("error in return")
    }
    if i != 1{
        fmt.Println("error in return")
    }
    if j != 1{
        fmt.Println("error in return")
    }
    if k != 1{
        fmt.Println("error in return")
    }
    if l != 1{
        fmt.Println("error in return")
    }
    if m != 1{
        fmt.Println("error in return")
    }
    if n != 1{
        fmt.Println("error in return")
    }
    if o != 1{
        fmt.Println("error in return")
    }
    if p != 1{
        fmt.Println("error in return")
    }
    if q != 1{
        fmt.Println("error in return")
    }
    if r != 1{
        fmt.Println("error in return")
    }
    if s != 1{
        fmt.Println("error in return")
    }
    */
}
