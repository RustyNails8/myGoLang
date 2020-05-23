package main

import(
  "fmt"
  "time"
)

func main() {
  now :=time.Now()
  year :=now.Year()
  fmt.Println(year)
  fmt.Print(now)
  // fmt.Print(year)
}
