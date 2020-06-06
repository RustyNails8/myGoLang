//Pckage reflectUsage uses reflect modulepackage main

import(
  "fmt"
  "reflect"
)

func main() {
  fmt.Println("Let me show type of 42    " , reflect.TypeOf(42))
  fmt.Println("Let me show type of 42.56 " , reflect.TypeOf(42.56))
  fmt.Println("Let me show type of Sumit " , reflect.TypeOf("Sumit"))
  fmt.Println("Let me show type of TRUE  " , reflect.TypeOf(true))

}
