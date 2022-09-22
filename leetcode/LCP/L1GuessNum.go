package main

func game(guess []int, answer []int) int {
	num := 0
	for i := 0; i < 3; i++ {
		if guess[i] == answer[i] {
			num++
		}
	}
	return num
}

func main() {}
