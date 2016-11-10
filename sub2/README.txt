Team Members: Hall, Pullig

Language: go

Problem: K- Guess the language

To Run/Compile:

  go run which_language.go file.txt
    -This will read in a file to determine what language is being used within.

Output: Based off of dictionary files that we were given our program will make a
        guess at what language was inputed. If multiple languages have the same
        number of words in their dictionary as others then all possible languages
        will be returned.

General Purpose Go Knowledge:
  General -> Go is a language built by Google that was meant to be as fast as C,
  as easy to develop in as Python, and make concurrency easy.

  layout of code ->
    a) every non-library executable file must start with package main
    b) afterwards you must import your libraries with import{ lib_name }
    c) you can optionally declare global variables between import and main
    c) all functions written between importing libraries/global vars and main
    d) you declare main

  function declarations -> each function declared as follow
    func name_of_function(name_of_var type_of_var) return_type{code;}

  variable declaration -> variables can be declared in both python and C style
    python style -> count := 5
    C style -> var count int

  Go Routines -> The single most impressive aspect of go is it's built in go
    routines. This is go's method of implementing multi-threaded programming in
    such an easy way that everyone will begin using it. To spawn a new thread/
    run a new go routine all you need to do is write go in front of a function
    call. This will cause the program to spawn a thread that will run that
    specific function and then die.

  Channel variables -> Another go specific resource is the channel variable.
    Channel variables are meant to be pipes that lead from the main of a program
    to its spawned threads, and then back to main. This is how you return data
    from a go routine and also how you force main to wait on the routines to
    finish.


Inner Working of our Code:

Used libraries and their purpose-
  fmt -> A formatting library used for output to a terminal screen
  os -> An OS library used to open and close files
  bufio -> A library used to read in a word file as strings instead of raw data
  io/ioutil -> A library used for directory manipulation and discovery
  strings -> A library used for manipulating strings

Functions-
  func how_many_words(c chan [2]int,language string,word string,which_lang int)
    This function takes in a channel int array of size 2 first. The second var is
    a string of which language  the dictionary is. The third var is a string of
    the phrase that we will be making our guesses about. The last variable is an
    int that will represent which language is being used.

    in this function we open up the dictionary that is dictated by the language
    string, loop through the dictionary check every word it contains by all of
    the words in the phrase string and keep a running count, then set the Channel
    int to be [count, language], and then return the channel int.

Main-
  our main reads in the phrase that we were given to analyze, looks up all of
  our dictionaries that we will be using, starts all of our go routines,
  collects the results of the go routines, finds the max count value of all of
  the go routines, and then prints out the names of all of the languages that
  the phrase could possibly be.
