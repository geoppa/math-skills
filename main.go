package main

import "os"

func main() {
	if len(os.Args) < 2 {
		println("Error. No data file declared")
		println("Usage: go run main.go <data_file>")
		return
	}
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		println("Error in file", err)
		return
	}
	print(file)
}
