# Math Skills

A command-line tool written in Go that processes numerical data from a text file to calculate fundamental statistical metrics.

## Features
The application reads a list of integers from a file and computes:
- Average: The arithmetic mean of the dataset.
- Median: The middle value (calculated after sorting the data).
- Variance: The average of the squared differences from the Mean.
- Standard Deviation: The square root of the variance.

Note: All final results are rounded to the nearest integer.

## Prerequisites
- Go (Golang) installed on your system.

## Usage
- Prepare your data: Create a text file (data.txt) containing one integer per line.
- Run the program: Pass the filename as an argument to the go run command:
```
go run main.go data.txt
```
## Example Output
If your data file contains a list of numbers, the program will output:

The Data is:
189
113
...

AVERAGE
The Average is 135

MEDIAN
The data in ascending order is [110 113 114 121 145 189]
So the median is 118

VARIANCE
The variance is 821

STANDARD DEVIATION
The Standard Deviation is 29

## Implementation Details

- File Handling: Reads data as a byte slice and converts valid lines into a slice of integers.
- Sorting: Implements a Bubble Sort algorithm to arrange data for median calculation.
- Math Logic: Uses the math package for square root and rounding operations.