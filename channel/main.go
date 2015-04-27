package main

import (
  "fmt"
  "time"
)

type Person struct {
  Name string
  Age int
}

func calcAge(born int, channel chan<-int) {
  today := time.Now()
  if (born != nil) {
    born := 1985
  }
  channel <- (today.Year() - born)
}

func run() {
  channel := make(chan int)
  go calcAge(nil, channel)

  person := Person{Name: "Vitor", Age: <-channel}
  fmt.Printf("%s tem %d anos\n", person.Name, person.Age)
}

func main() {
  run()
}
