package main

import "fmt"

func main() {
	var myName = "Bora"
	fmt.Println("my name is", myName, myName)

	var name string = "Kathy"
	fmt.Println("name =", name)

	userName := "admin"
	fmt.Println("userName =", userName)

	var sum int
	fmt.Println("The sum is =", sum)

	part1, other := 1, 5
	fmt.Println("part1 =", part1, "other is", other)

	part2, other := 1, 5
	fmt.Println("part1 =", part2, "other is", other)

	sum = part1 + part2
	fmt.Println("sum= ", sum)

	var (
		lessonName = "Varaibles"
		lessontype = "Demo"
	)

	fmt.Println("lessonName", lessonName)
	fmt.Println("lessonType", lessontype)

	word1, word2, _ := "hello", "world", "!"
	fmt.Println(word1, word2)

}
