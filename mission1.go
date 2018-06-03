package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func do_calc_loop(times int, quit chan int) {
	reader := bufio.NewReader(os.Stdin)
	if times <= 0 {
		quit <- 0
	} else {
		var length int
		fmt.Scanln(&length)
		line, _ := reader.ReadString('\n')
		//fmt.Println(line)
		lines := strings.Fields(line)
		//fmt.Println(lines)
		go calc_one_slice(length, 0, 0, lines)
		do_calc_loop(times-1, quit)
	}
}

func calc_one_slice(length int, index int, total int, s []string) {
	if index > length-1 {
		fmt.Println(total)
	} else {
		num, _ := strconv.Atoi(s[index])
		if num < 0 {
		} else {
			total += num * num
		}
		calc_one_slice(length, index+1, total, s)
	}
}

func main() {
	var num int
	fmt.Scanln(&num)
	quit := make(chan int)
	go do_calc_loop(num, quit)
	num = <-quit
}
