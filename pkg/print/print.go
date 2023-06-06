package print

// Instead of these:
func printStrings(s []string) {
	//loop and print
}
func printIntegers(s []int) {
	//loop and print
}

// we can do this:
func printAnything(s interface{}) {
	//loop and print
}

func testFunction() {
	strings := []string{"apple", "orange", "banana"}
	integers := []int{1, 2, 3}
	printAnything(strings)
	printAnything(integers)
	// Note that the empty interface/any type should be used carefully only after the developer has considered possible type errors that could occur with the function.
}
