package main

func CountLights() int {

	countOn := 0

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			countOn += Matrix[i][j]
		}
	}
	return countOn
}
