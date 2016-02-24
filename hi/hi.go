package main

import {
 "bufio"
 "os"
 "fmt"
}

func main() {
  r := bufio.NewReader(os.Stdin)

  println("What is your name?")
  line, err := ReadString('\n')
  if err != nil {
    println(err)
    os.Exit(1)
  }
  fmt.Print("Hello" + line)

)
