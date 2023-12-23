// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	numberA := 0
	numberB := 1
	
	for i:=0; i <= 100; i++{
		fibonacci := numberA + numberB
		if fibonacci >= 100000 {
			break
		}
		numberA = numberB
		numberB = fibonacci
		fmt.Println("Index:  ---",i,"Fibonacci: ",fibonacci)
	}
			
}