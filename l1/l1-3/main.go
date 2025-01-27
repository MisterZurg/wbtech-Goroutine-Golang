package main

import "fmt"

var array = []int64{2, 4, 6, 8, 10}

func sumOfSquares(nums []int64, sq chan int64, done chan struct{}) {
	for _, n := range nums {
		sq <- n * n
	}

	done <- struct{}{}
}

func main() {
	sum := int64(0)
	// За водим каналы, для передачи квадратов и сигнала завершения
	squares, done := make(chan int64), make(chan struct{})

	// Запускаем горутину, которая фоново будет считать квадраты
	go sumOfSquares(array, squares, done)

	// Запускаем горутину, которая дождётся выполнения /\
	// и закроет канал результата
	go func() {
		<-done
		close(squares)
	}()

	for v := range squares {
		sum += v
	}
	fmt.Println(sum)
}
