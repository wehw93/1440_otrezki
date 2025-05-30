package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var n int
	
	file, err := os.Open("data_prog_contest_problem_1.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		fmt.Println("Файл пуст")
		os.Exit(1)
	}
	n, err = strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Ошибка чтения числа n:", err)
		os.Exit(1)
	}
	otrezki := make([][2]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		str := scanner.Text()
		numbers := strings.Fields(str)
		x, _ := strconv.Atoi(numbers[0])
		y, _ := strconv.Atoi(numbers[1])
		otrezki[i] = [2]int{x, y}
	}

	/*
		fmt.Scan(&n)

		scanner:=bufio.NewScanner(os.Stdin)

		for i:=0; i < n; i++{
			scanner.Scan()
			str:=scanner.Text()
			numbers:=strings.Fields(str)
			x,_:=strconv.Atoi(numbers[0])
			y,_:=strconv.Atoi(numbers[1])
			otrezki[i] = [2]int{x,y}
		}
	*/
	res := Foo(otrezki)
	fmt.Printf("%d\n", res)

}

func Foo(otrezki [][2]int) int {
	if len(otrezki) == 0 {
		return 0
	}
	if len(otrezki) == 1 {
		return 1
	}
	sort.Slice(otrezki, func(i, j int) bool {
		return otrezki[i][1] < otrezki[j][1]
	})
	res := 0
	mostRight := -1
	for _, otrezok := range otrezki {
		if mostRight < otrezok[0] {
			mostRight = otrezok[1]
			res++
		}
	}
	return res
}
