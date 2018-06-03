// HDE Challenge 003
// Mission 1
// Zhongfan Dou 2018/6

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Do the calculation loop, send signal to quit channel if reaches end
func do_calc_loop(times int, quit chan int) {
	reader := bufio.NewReader(os.Stdin)
	if times <= 0 {
		quit <- 0
	} else {
		var length int
		_, err := fmt.Scanln(&length)
		if err {
			fmt.Println("Not reading one integer")
			os.Exit(1)
		}
		line, _ := reader.ReadString('\n')
		nums := strings.Fields(line)
		go calc_one_slice(len(nums), 0, 0, nums)
		do_calc_loop(times-1, quit)
	}
}

// For each test case, calculate the sum of squares of non-negative num and print
func calc_one_slice(length int, index int, total int, s []string) {
	if index > length-1 {
		fmt.Println(total)
	} else {
		num, err := strconv.Atoi(s[index])
		if err {
			fmt.Println("Input is not valid integer")
			os.Exit(1)
		} else if num < 0 || n > 100 {
			// do not calc for number not in range
		} else {
			total += num * num
		}
		calc_one_slice(length, index+1, total, s)
	}
}

func main() {
	var num int
	_, err := fmt.Scanln(&num)
	if err {
		fmt.Println("first line must be one integer of how many test cases")
		os.Exit(1)
	} else if num < 0 || num > 100 {
		fmt.Println("Number of test cases must be between 0 and 100")
		os.Exit(1)
	}
	quit := make(chan int)
	go do_calc_loop(num, quit)
	num = <-quit
}
