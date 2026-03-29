package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	// check if at least one file passed as argument
	if len(os.Args) < 2 {
		println("Error. No data file declared \nUsage: go run main.go <data_file>")
		return
	}
	// read the name of the file and read the contents (in byte slice)
	filename := os.Args[1]
	data, err := os.ReadFile(filename)
	// check if the filename given as argument exists
	if err != nil {
		fmt.Printf("Error:%v\n", err.Error())
		return
	}
	// check if file is empty
	if len(data) == 0 {
		fmt.Println("Empty file")
		return
	}
	// find total lines, split numbers by line, calc sum, make array of int
	lines := bytes.Split(data, []byte("\n"))
	var nums []int
	sum := 0
	total := 0
	for _, line := range lines {
		num, err := strconv.Atoi(string(line))
		if err == nil {
			nums = append(nums, num)
			sum += num
			total++
		}
	}
	// print the data the sum and the total numbers
	fmt.Printf("The Data is:\n%s\n", data)
	fmt.Printf("Total Numbers:%v\n", total)
	fmt.Printf("Sum:%v\n\n", sum)
	// find and print the average
	average := float64(sum) / float64(total)
	fmt.Printf("AVERAGE\nThe Average is %.2f\n\n", average)
	// bubble sort accending the array of int
	for i := 0; i < total-1; i++ {
		for j := 0; j < total-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	// find the median
	var med float64
	if total%2 == 1 {
		med = float64(nums[total/2])
	} else {
		med = float64((nums[total/2-1] + nums[total/2])) / 2
	}
	fmt.Println("MEDIAN\nThe data in accending order is", nums)
	fmt.Printf("So the median is %.2f\n\n", med)
	// find the variance and the standard deviation
	var variSum float64 = 0
	var diff float64 = 0
	var vari float64 = 0
	for i := 0; i < total; i++ {
		diff = float64(nums[i]) - float64(average)
		variSum = variSum + (diff * diff)
	}
	vari = variSum / float64(total)
	stddev := math.Sqrt(vari)
	fmt.Printf("VARIANCE\nThe variance is %.2f\n\n", vari)
	fmt.Printf("STANDARD DEVIATION\nThe Standard Deviation is %.2f\n", stddev)
}
