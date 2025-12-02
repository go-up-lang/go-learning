package main

import "fmt"

//// 8
type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "ramen"}
	nico := person{name: "nico", age: 18, favFood: favFood}
	fmt.Println(nico)
}

//// 1. function print
// func lenAn(name string) (int, string) {
// 	return len(name), string(name)
// }

// func main() {
// 	totalLenght, upperName := lenAn("nico")
// 	fmt.Println(totalLenght, upperName)
// }
//// --------------------------------------------------

//// 2. function print
// func lenAn(name string) (lenght int, uppercase string) {
// 	lenght = len(name)
// 	uppercase = strings.ToUpper(name)
// 	return
// }

// func main() {
// 	totalLenght, up := lenAn("nico")
// 	fmt.Println(totalLenght, up)
// }
//// --------------------------------------------------

//// 3. function defer
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
//// --------------------------------------------------

//// 4. function for
// func addFor(numbers ...int) int {
// 	fmt.Println(numbers)
// 	for number, index := range numbers {
// 		fmt.Println(number, index)
// 	}
// 	return 0
// }

// func main() {
// 	addFor(1, 2, 3, 4, 5, 6)
// }
//// --------------------------------------------------

//// 4. function if~else
// func canIDrink(age int) bool {
// 	if koreanAge := age + 2; koreanAge < 18 {
// 		return false
// 	}
// 	return true
// }

// func main() {
// 	fmt.Println(canIDrink(17))
// }
//// --------------------------------------------------

//// 5. function point
// func main() {
// 	a := 2
// 	b := &a
// 	a = 5
// 	fmt.Println(a, *b)
// }
//// --------------------------------------------------

//// 6. function array
// func main() {
// 	names := [5]string{"nico", "lynn", "gyusun"}
// 	names[2] = "yayal"
// 	names[4] = "a;a;a;"
// 	// names[5] = "ohoho"
// 	fmt.Println(names)
// }
//// --------------------------------------------------

//// 7. function maps
// func main() {
// 	nico := map[string]string{"name": "nico", "age": "12"}
// 	for _, value := range nico {
// 		fmt.Println(value)
// 	}
// }
//// --------------------------------------------------
