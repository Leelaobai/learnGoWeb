package main

import (
	"C2"
	"fmt"
)

type testInt func(int) bool

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func testFilter() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("slice = ", slice)
	odd := filter(slice, C2.IsOdd)
	fmt.Println("Odd elements of slice are: ", odd)
	even := filter(slice, C2.IsEven)
	fmt.Println("Even elements of slice are: ", even)
}

func testStruct() {
	var tom C2.Person
	tom.Name, tom.Age = "Tom", 18

	bob := C2.Person{Age: 25, Name: "Bob"}

	paul := C2.Person{"Paul", 43}

	tb_Older, tb_diff := C2.Older(tom, bob)
	tp_Older, tp_diff := C2.Older(tom, paul)
	bp_Older, bp_diff := C2.Older(bob, paul)
	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.Name, bob.Name, tb_Older.Name, tb_diff)
	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.Name, paul.Name, tp_Older.Name, tp_diff)
	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		bob.Name, paul.Name, bp_Older.Name, bp_diff)

}

func main() {
	//C2.HelloC2()
	//C2.Define_const()
	//C2.Test_boolean()
	//C2.Test_number()
	//C2.Test_string()
	//C2.Test_iota()
	//C2.Test_array()
	//C2.Test_ifelse()
	//C2.Test_range_map()
	//testStruct()
	C2.TestMethod()
}