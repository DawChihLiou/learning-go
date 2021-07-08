package main

import (
  "fmt"
  "log"
  "example.com/greetings"
)

func main() {
  // set properties of the predefined Logger, including
  // the log entry prefix and a flag to disable priting 
  // the time, source file, and line number.
  log.SetPrefix("greetings: ")

  log.SetFlags(0)

  // request a greeting message
  message, err := greetings.Hello("Gladys")
  // If an error was returned, print it to the concole
  // and exit the program.
  if err != nil {
    log.Fatal(err)
  }
  // if no error was returned, print the returned message
  // to the concole
  fmt.Println(message)
}
