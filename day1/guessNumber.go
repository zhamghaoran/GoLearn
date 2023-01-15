package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano()) // 初始化随机数种子
	secretNum := rand.Intn(maxNum)
	fmt.Println("please input your number")
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n') // 读取输入
		if err != nil {
			fmt.Println("An error occured while reading inout. Please try again")
			continue
		}
		input = strings.Trim(input, "\r\n") // 去掉换行符
		guess, err := strconv.Atoi(input)   // 转化成数字
		if err != nil {
			fmt.Println("Invaild input. Please enter an integer value")
			continue
		}
		fmt.Println("Your guess number is ", guess)
		if guess > secretNum {
			fmt.Println("Your guess number is bigger than the secret number. Please try again")
		} else if guess < secretNum {
			fmt.Println("Your guess number is smaller than the secret number. Please try again")
		} else {
			fmt.Println("correct")
			break
		}
	}

}
