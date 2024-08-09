package main //if the package is main, it tells the GO compiler this code should be compiled into executable program at the end

import "fmt"

func main() {
	// var ages [3]int = [3]int{23, 34, 53}
	// fmt.Println(ages)

	// var ages = [3]int{30, 34, 21}
	names := [4]string{"Mary", "Bill", "Lily", "Kate"}
	// names[2] = "Eason"
	// fmt.Println(ages, len(ages))
	// fmt.Println(names, len(names))

	//slices (use arrays under the hood)
	// var scores = []int{100, 98, 97, 56}
	// scores[1] = 99
	// scores = append(scores, 78)
	// fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

	rangeOne = append(rangeOne, "Helen")
	fmt.Println(rangeOne)

}
