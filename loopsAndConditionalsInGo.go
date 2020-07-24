package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	for i := 0; i <= 10; i++ {
		fmt.Printf("The outer loop: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\tThe inner loop: %d\n", j)
		}
	}

	// There is no while loop in Go. Instead use for with just a conditional statement
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("done")

	if true {
		fmt.Println("This will always print")
	}

	// Using the NOT operator
	if !false {
		fmt.Println("This too will always print")
	}

	if false {
		fmt.Println("However, this will never print")
	}

	// You can also use an if statement in an infinite loop to make a while loop
	y := 1
	for {
		if y > 9 {
			break
		}
		fmt.Println(y)
		y++
	}
	fmt.Println("done")

	// Print all the odd numbers between 0 and 100 using an infinit loop and utilizing break and continue
	z := 0
	for {
		if z > 100 {
			break
		}
		if z%2 == 0 {
			z++
			continue
		}
		fmt.Println(z)
		z++
	}
	fmt.Println("done")

	// Printing out all of the ascii characters that correspond to the numbers between 33 and 122 included
	for i := 33; i <= 122; i++ {
		fmt.Printf("decimal: %v\tASCII: %#U\n", i, i)
	}

	// Conditional with an initialization statement
	// Helps to limit the declared variable's scope
	if x := 42; x == 2 {
		fmt.Println("This will not print")
	}

	// if - esle if - else
	a := 42
	if a == 40 {
		fmt.Println("a is 40")
	} else if a == 41 {
		fmt.Println("a is 41")
	} else {
		fmt.Println("a is not 40 or 41")
	}
}
