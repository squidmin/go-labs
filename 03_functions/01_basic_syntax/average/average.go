package main

func average(numbers []float64) (average float64) {
	total := 0.0
	for _, v := range numbers {
		total += v
	}
	average = total / float64(len(numbers))
	return
}

// Without using named return type:
//func average(numbers []float64) float64 {
//	total := 0.0
//	for _, v := range numbers {
//		total += v
//	}
//	average := total / float64(len(numbers))
//	return average
//}

func main() {
	average([]float64{1, 2, 3, 4, 5})
}
