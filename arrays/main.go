package main

import "fmt"

func main() {
	notes := [3]string{"DO", "RE", "MI"}
	primes := [5]int{2, 3, 5, 7, 11}

	fmt.Println(notes)
	fmt.Println(primes)
	fmt.Printf("%#v\n", notes)
	fmt.Printf("%#v\n", primes)

	index := 3
	fmt.Println(notes[index])

}
