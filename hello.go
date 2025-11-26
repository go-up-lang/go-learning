package main

import "fmt"

func addFor(numbers ...int) int {
	fmt.Println(numbers)
	for number, index := range numbers {
		fmt.Println(number, index)
	}
	return 0
}

func main() {
	addFor(1, 2, 3, 4, 5, 6)
}

//// 1. function1 -------------------------
// func lenAn(name string) (int, string) {
// 	return len(name), string(name)
// }

// func main() {
// 	totalLenght, upperName := lenAn("nico")
// 	fmt.Println(totalLenght, upperName)
// }

//// 2. function2 -------------------------
// func lenAn(name string) (lenght int, uppercase string) {
// 	lenght = len(name)
// 	uppercase = strings.ToUpper(name)
// 	return
// }

// func main() {
// 	totalLenght, up := lenAn("nico")
// 	fmt.Println(totalLenght, up)
// }

// // 3. function defer -------------------------
// func lenAn(name string) (lenght int, uppercase string) {
// 	defer fmt.Println("this is defer")
// 	lenght = len(name)
// 	uppercase = strings.ToUpper(name)
// 	return
// }

// func main() {
// 	totalLenght, up := lenAn("nico")
// 	fmt.Println(totalLenght, up)
// }
