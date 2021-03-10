package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var students []int

func main() {
	// 从标准输入流中接收输入数据
	input := bufio.NewScanner(os.Stdin)

	fmt.Printf("Please type in something:\n")

	result := make([]int, 0)
	lineNum := 0
	maxLineNum := 2
	// 逐行扫描
	for input.Scan() {
		line := input.Text()
		switch lineNum {
		case 0:
			ss := strings.Split(line, " ")
			stuNum, _ := strconv.Atoi(ss[0])
			optNum, _ := strconv.Atoi(ss[1])
			students = make([]int, stuNum)
			maxLineNum += optNum
		case 1:
			ss := strings.Split(line, " ")
			ss = ss[:len(students)]
			for i, v := range ss {
				students[i], _ = strconv.Atoi(v)
			}
		default:
			ss := strings.Split(line, " ")
			ss = ss[:3]
			switch ss[0] {
			case "Q":
				index1, _ := strconv.Atoi(ss[1])
				index2, _ := strconv.Atoi(ss[2])
				index1 -= 1
				result = append(result, max(students[index1:index2]))
			case "U":
				index, _ := strconv.Atoi(ss[1])
				val, _ := strconv.Atoi(ss[2])
				index -= 1
				students[index] = val
			}

		}
		// 输入bye时 结束
		lineNum++
		if lineNum >= maxLineNum {
			break
		}
		if line == "bye" {
			break
		}

	}

	for _, v := range result {
		fmt.Println(v)
	}

}

func max(in []int) int {
	max := in[0]
	for i := 1; i < len(in); i++ {
		if in[i] > max {
			max = in[i]
		}
	}
	return max
}
