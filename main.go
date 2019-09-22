package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFile(filename string, separator func(data []byte, atEOF bool) (advance int, token []byte, err error), integerChan chan int, signalChan chan int){
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		fmt.Print(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(separator)
	var sum int
	for scanner.Scan() {

		value, err := strconv.Atoi(scanner.Text())
		if err == nil {
			sum += value
		}
	}
	integerChan <- sum
	signalChan <- 1
}

func main(){
	//start := time.Now()
	separatorFunc := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' || data[i] == '\r' || data[i] == '\n' {
				return i + 1, data[:i], nil
			}
		}
		return 0, data, bufio.ErrFinalToken
	}

	integerChan, signalChan := make(chan int), make(chan int)
	for dirIndex := 1; dirIndex <= 991; dirIndex += 10 {
		var basePath string
		if dirIndex == 91 {
			basePath = fmt.Sprintf("files/0000%d-000%d/", dirIndex, dirIndex + 9)
		} else if dirIndex == 991 {
			basePath = fmt.Sprintf("files/000%d-00%d/", dirIndex, dirIndex + 9)
		} else if dirIndex > 91 {
			basePath = fmt.Sprintf("files/000%d-000%d/", dirIndex, dirIndex + 9)
		} else if dirIndex > 1 {
			basePath = fmt.Sprintf("files/0000%d-0000%d/", dirIndex, dirIndex + 9)
		} else {
			basePath = fmt.Sprintf("files/00000%d-0000%d/", dirIndex, dirIndex + 9)
		}

		for index := dirIndex; index <= dirIndex + 9; index++ {

			var path string
			if index >= 1000{
				path = fmt.Sprintf("%s00%d.csv", basePath, index)
			} else if index >= 100 {
				path = fmt.Sprintf("%s000%d.csv", basePath, index)
			} else if index >= 10 {
				path = fmt.Sprintf("%s0000%d.csv", basePath, index)
			} else {
				path = fmt.Sprintf("%s00000%d.csv", basePath, index)
			}
			go readFile(path, separatorFunc, integerChan, signalChan)
		}

	}

	var sum, signal int
	for ; signal < 1000 ; signal += <- signalChan {
		sum += <- integerChan
	}

	//finished := time.Now()
	//elapsed := finished.Sub(start)
	//
	//fmt.Printf("The sum is %d gotten after %s", sum, fmt.Sprint(elapsed))


}
