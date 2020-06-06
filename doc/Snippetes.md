###Snippets

1. Go MAIN
```
package main

import (
  "fmt"
  )

func main() {

}
```
2. USER INPUT
```
	fmt.Print("Enter the score:   ")
	// reader := bufio.NewReader(os.Stdin)
	// LPAR, err := reader.ReadString('\n')
	score, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	score = strings.TrimSpace(score)

	grade, err := strconv.ParseFloat(score, 64)
	if err != nil {
		fmt.Println("\n\nWould you please enter a number ?")
		fmt.Println("Please rerun the program !\n\n I got this error : ")
		log.Fatal(err)
	}
```

3. FOR LOOP
```
	for tries := 0; tries < 10; tries++ {
}
```

4. IF
```
		if err != nil {
			log.Fatal(err)
		}
```
```
		if guess < target {
			fmt.Println(" Too low  a guess ...")
		} else if guess > target {
			fmt.Println(" Too high  a guess ...")
		} else if guess == target {
			success = true
			fmt.Println(" Correct    guess ...")
			break
		}
	}
```
```
	if !success {
		fmt.Println(" The target was ... ", target)
	}
```