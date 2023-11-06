package main

func main() {

	//slice of integers from 1 to 10
	num := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < len(num); i++ {

		if i%2 == 0 {
			println(i, "is even")
		} else {
			println(i, "is odd")
		}

	}

}
