package main

import(
  "bufio"
  "fmt"
  "os"
)

func main() {
  fmt.Print("Enter Server Name:   ")
  // Server := bufio.NewReader(os.Stdin)
  // input := Server.ReadString('\n')
  LPAR, err :=bufio.NewReader(os.Stdin).ReadString('\n')
  fmt.Println(LPAR, err)
}
