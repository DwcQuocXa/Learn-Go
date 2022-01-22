package main

import (
	"errors"
	"fmt"
	"math"
)

 type person struct {
 	name string
 	age int
 }

func main () {

	 p := person{name:"Duc", age: 20}
	 fmt.Println(p)
	 fmt.Println(p.age)

	 // a := [5]int {5,4,3,2,1}
	 a := []int {5,4,3,2,1}
	 a = append(a, 13)

	  fmt.Println("Hello Golang")
	  fmt.Println(Sum(3))
	 fmt.Println(a)

	  vertices := make(map[string]int)

	  vertices["triangle"] = 2
	  vertices["square"] = 3
	  vertices["dodecagon"] = 12

	  delete(vertices, "square")

	  fmt.Println(vertices)

	 arr := []string{"a", "b", "c"}

	 for index, value := range arr {
	 	fmt.Println("index:", index, "value:", value)
	 }

	 for key, value := range vertices {
	 	fmt.Println("key:", key, "value:", value)
	 }
	 result, err := sqrt(-16)

	 if err != nil {
	 	fmt.Println(err)
	 } else {
	 	fmt.Println(result)
	 }
	
	i := 7
	fmt.Println(i)
	fmt.Println(&i)
	inc(&i)
	fmt.Println(i)
}

func inc (x *int) {
	*x++
}

 func Sum (number int) (int) {
 	sum := 0
 	for i:=0; i<number; i++ {
 		sum +=i
 	}
 	return sum
 }

 func sqrt (x float64) (float64, error) {
 	if x <0 {
 		return 0, errors.New("Undefined or negative num")
 	}

 	return math.Sqrt(x), nil
 }

